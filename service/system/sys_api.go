package system

import (
	"errors"
	"fmt"
	"server/global"
	"server/model/common/request"
	"server/model/system"

	"gorm.io/gorm"
)

type ApiService struct{}

// 创建api
func (apiService *ApiService) CreateApi(api system.SysApi) error {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的api")
	}
	return global.GVA_DB.Create(&api).Error
}

// 删除api
func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	var entity system.SysApi
	fmt.Println(api.ID)
	err = global.GVA_DB.Where("id = ?", api.ID).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	//同步清理casbin
	success := CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	if !success {
		return errors.New(entity.Path + ":" + entity.Method + "casbin同步清理失败")
	}
	e := CasbinServiceApp.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

func (apiService *ApiService) GetApiById(id int) (api system.SysApi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

// 更新对应api信息
func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.GVA_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = global.GVA_DB.Save(api).Error
	}
	return err
}

// 删除对应api,多个删除
func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	var apis []system.SysApi
	err = global.GVA_DB.Find(&apis, "id in ?", ids.Ids).Delete(&apis).Error
	if err != nil {
		return err
	} else { //
		for _, sysApi := range apis {
			success := CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
			if !success {
				return errors.New(sysApi.Path + ":" + sysApi.Method + "casbin同步清理失败")
			}
		}
		e := CasbinServiceApp.Casbin()
		err = e.InvalidateCache()
		if err != nil {
			return err
		}
	}

	return
}

// 获得所有api
func (apiService *ApiService) GetAllApis() (apis []system.SysApi, err error) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

// 获取对应api
func (apiService *ApiService) GetAPIInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysApi{})
	var apiList []system.SysApi
	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}
	err = db.Count(&total).Error
	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else {
				err = fmt.Errorf("非法的排序字段: %v", order)
				return apiList, total, err
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return apiList, total, err
}
