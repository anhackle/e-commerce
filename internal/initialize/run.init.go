package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitMysqlC()
	InitRedis()
	InitServiceInterface()
	InitValidator()

	r := InitRouter()
	return r

	// r.Run(":8082")
}
