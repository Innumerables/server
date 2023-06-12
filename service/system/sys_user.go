package system

import (
	"errors"
	"fmt"
	"server/global"
	"server/model/common/request"
	"server/model/system"
	"time"

	"server/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserService struct {
}

// 用户登录
func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user system.SysUser
	// err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").
	// 	First(&user).Error
	err = global.GVA_DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		// MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

// 注册用户
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

// 改变用户密码
func (userService *UserService) ChangePasword(u *system.SysUser, newPassword string) (err error) {
	var user system.SysUser
	err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error
	if err != nil {
		return err
	}
	ok := utils.BcryptCheck(u.Password, user.Password)
	if !ok {
		return errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return err
}

// 也是设置用户权限，有所缺陷，应该被舍去了
func (userService *UserService) SetUserAuthotity(id uint, authorithId uint) (err error) {
	assingErr := global.GVA_DB.Where("sys_user_id = ? AND sys_autority_authority_id = ?", id, authorithId).First(&system.SysAuthority{}).Error
	if errors.Is(assingErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GVA_DB.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorithId).Error
	return err
}

// 删除用户
func (userService *UserService) DeleteUser(id int) (err error) {
	var user system.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
	return err
}

// 设置用户信息
func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return global.GVA_DB.Model(&system.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "sideMode", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"side_mode":  req.SideMode,
			"enable":     req.Enable,
		}).Error
}

// 在设置用户权限时，会从数据库中删除所有相关的权限,通过前端获得的权限id再更新到数据库中
func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		return nil
	})
}

// 更新用户信息
func (userService *UserService) SetSelfInfo(req system.SysUser) error {
	return global.GVA_DB.Model(&system.SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}

// 重置密码
func (userService *UserService) ResetPassword(id uint) error {
	err := global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", id).Update("password", utils.BcryptHash("123456")).Error
	return err
}

// 根据设定的页数和取数据数获取对应的用户信息
func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

// 根据id获得对应用户信息
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.GVA_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	// MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}
