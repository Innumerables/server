package system

import "server/service"

type ApiGroup struct {
	DBApi
	BaseApi
	SystemApiApi
	CasbinApi
	AuthorityApi
	AuthorityMenuApi
}

var (
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService    = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	apiService       = service.ServiceGroupApp.SystemServiceGroup.ApiService
	casbinService    = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	menuService      = service.ServiceGroupApp.SystemServiceGroup.MenuService
	baseMenuService  = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
)
