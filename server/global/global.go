package global

import (
	"go.uber.org/zap"

	"server/config"
)

//全局对象包

var (
	Config *config.Config
	Log    *zap.Logger
)
