package initialize

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
