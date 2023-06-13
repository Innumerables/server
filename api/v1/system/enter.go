package system

import "server/service"

type ApiGroup struct {
	DBApi
	BaseApi
	SystemApiApi
	CasbinApi
	AuthorityApi
}

var (
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService    = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	apiService       = service.ServiceGroupApp.SystemServiceGroup.ApiService
	casbinService    = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthroityService
)
