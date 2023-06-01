package main

import (
	"server/global"

	"server/core"

	"server/initialize"

	"go.uber.org/zap"
)

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() //初始化viper
	//initialize.OtherInit()
	global.GVA_LOG = core.Zap() //初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() //gorm连接数据库
	// initialize.Timer()
	// initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables()
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServe()
}
