package initialize

import (
	"os"
	"server/global"

	"server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	// case "pgsql":
	// 	return GormPgSql()
	// case "oracle":
	// 	return GormOracle()
	// case "mssql":
	// 	return GormMssql()
	default:
		return GormMysql()
	}
}

// 注册数据库表专用，自动迁移
func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		// 系统模块表
		system.SysUser{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
