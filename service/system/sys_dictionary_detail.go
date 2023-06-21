package system

import (
	"server/global"
	"server/model/system"
	"server/model/system/request"
)

type DictionaryDetailService struct{}

func (detail *DictionaryDetailService) CreateSysDictionaryDetail(dictionDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Create(&dictionDetail).Error
	return
}

func (detail *DictionaryDetailService) DeleteSysDictionaryDetail(dictionDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Delete(&dictionDetail).Error
	return
}

func (detail *DictionaryDetailService) UpadateSysdictionaryDetail(dicitionDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Save(&dicitionDetail).Error
	return
}

func (detail *DictionaryDetailService) FindSysDictionaryDetail(id uint) (sysDictionaryDetail system.SysDictionaryDetail, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

func (detail *DictionaryDetailService) GetSysDictionaryDetailList(info request.SysDictionaryDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	//利用多个条件进行查询
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&sysDictionaryDetails).Error
	return sysDictionaryDetails, total, err
}
