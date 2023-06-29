package example

import (
	"server/global"
	"server/model/common/request"
	"server/model/example"
	"server/model/system"
	systemService "server/service/system"
)

type CustomerService struct{}

func (customerService *CustomerService) CreateCustomer(customer example.Customer) error {
	err := global.GVA_DB.Create(&customer).Error
	return err
}

func (customerService *CustomerService) UpdateCustomer(customer *example.Customer) error {
	err := global.GVA_DB.Save(customer).Error
	return err
}

func (customerService *CustomerService) DeleteCustomer(customer example.Customer) error {
	err := global.GVA_DB.Delete(&customer).Error
	return err
}

func (customerService *CustomerService) GetCustomer(id uint) (customer example.Customer, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

func (customerService *CustomerService) GetCustomerList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&example.Customer{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []uint
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []example.Customer
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}
