package system

type RouterGroup struct {
	BaseRouter
	InitRouter
	ApiRouter
	CasbinRouter
	UserRouter
	AuthorityRouter
	MenuRouter
}
