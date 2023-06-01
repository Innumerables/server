package system

type ServiceGroup struct {
	UserService
	InitDBService
	CasbinService
	ApiService
}
