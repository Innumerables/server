package system

type ServiceGroup struct {
	UserService
	InitDBService
	CasbinService
	ApiService
	AuthorityService
	MenuService
	BaseMenuService
	DictionaryService
	DictionaryDetailService
	OperationRecordService
}
