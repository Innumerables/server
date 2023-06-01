package system

import (
	"errors"
	"fmt"
	"server/global"
	"server/model/system"

	"server/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserService struct {
}

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

func (userService *UserService) SetUserAuthotity(id uint, authorithId uint) (err error) {
	assingErr := global.GVA_DB.Where("sys_user_id = ? AND sys_autority_authority_id = ?", id, authorithId).First(&system.SysAuthority{}).Error
	if errors.Is(assingErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GVA_DB.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorithId).Error
	return err
}
