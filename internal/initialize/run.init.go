package initialize

import (
	"fmt"

	"github.com/anle/codebase/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
	InitValidator()

	r := InitRouter()
	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))

}
