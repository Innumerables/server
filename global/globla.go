package global

import (
	"server/config"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_DB                  *gorm.DB
	GVA_Concurrency_Control = &singleflight.Group{}
	BlackCache              local_cache.Cache
)
