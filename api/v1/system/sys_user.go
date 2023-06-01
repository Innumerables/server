package system

import (
	"server/global"
	"server/model/common/response"
	"server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"

	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	u := &system.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	b.TokenNext(c, *user)
	return
}

func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

}

func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	usesrReturn, err := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: usesrReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: usesrReturn}, "注册成功", c)
}

func (b *BaseApi) ChangePasword(c *gin.Context) {
	var cgp systemReq.ChangePasword
	err := c.ShouldBindJSON(&cgp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: cgp.Password}
	err = userService.ChangePasword(u, cgp.Password)
	if err != nil {
		global.GVA_LOG.Error("修改失败", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

func (b *BaseApi) SetUserAuthotity(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = userService.SetUserAuthotity(userId, sua.AuthorityId)
}

func (b *BaseApi) DeleteUser(c *gin.Context) {

}

func (b *BaseApi) SetuserInfo(c *gin.Context) {

}

func (b *BaseApi) SetSelfInfo(c *gin.Context) {

}

func (b *BaseApi) SetUserAuthorities(c *gin.Context) {

}

func (b *BaseApi) ResetPassword(c *gin.Context) {

}

func (b *BaseApi) GetUserList(c *gin.Context) {

}

func (b *BaseApi) GetUserInfo(c *gin.Context) {

}
