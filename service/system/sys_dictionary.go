package system

import (
	"errors"
	"server/global"
	"server/model/system"
	"server/model/system/request"

	"gorm.io/gorm"
)

type DictionaryService struct{}

func (dictionaryService *DictionaryService) CreateSysDictionary(dictionary system.SysDictionary) (err error) {
	if (!errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", dictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.GVA_DB.Create(&dictionary).Error
	return err
}

func (dictionaryService *DictionaryService) DeleteSysDictionary(dictionary system.SysDictionary) (err error) {
	err = global.GVA_DB.Where("id = ?", dictionary.ID).Preload("sysDictionaryDetails").First(&dictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("未发现此纪录")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&dictionary).Error
	if err != nil {
		return err
	}
	if dictionary.SysDictionaryDetails != nil {
		return global.GVA_DB.Where("sys_dictionary_id = ?", dictionary.ID).Delete(dictionary.SysDictionaryDetails).Error
	}
	return
}

func (dictionaryService *DictionaryService) UpdateSysDictionary(dictionary system.SysDictionary) (err error) {
	var dict system.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   dictionary.Name,
		"Type":   dictionary.Type,
		"Status": dictionary.Status,
		"Desc":   dictionary.Desc,
	}
	db := global.GVA_DB.Where("id = ?", dictionary.ID).First(&dict)
	if dict.Type != dictionary.Type {
		if !errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", dictionary.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的type，不允许更改")
		}
	}
	err = db.Updates(sysDictionaryMap).Error
	return err
}

// 此方法是在更改字典详情时使用，用来获得当前字典的信息加载到前端以便用户更改
func (dictionaryService *DictionaryService) GetSysDictionary(Type string, Id uint, status *bool) (sysDictionary system.SysDictionary, err error) {
	var flag = false
	if status == nil {
		flag = true
	} else {
		flag = *status
	}
	err = global.GVA_DB.Where("(type = ? OR id = ?) and status = ?", Type, Id, flag).Preload("SysDictionaryDetails", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", true).Order("sort")
	}).First(&sysDictionary).Error
	return
}

// 查询相关字典，并分页获取查到的数据
func (dictionaryService *DictionaryService) GetSysDictionaryInfoList(info request.SysDictionarySearch) (list interface{}, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysDictionary{})
	var sysDictionarys []system.SysDictionary
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sysDictionarys).Error
	return
}
