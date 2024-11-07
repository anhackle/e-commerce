package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/anle/codebase/global"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)

	db, err := sql.Open("mysql", s)

	checkErrorPanicC(err, "InitMysql initialization error")

	global.Logger.Info("Initializing MySQL Successfully")

	global.Mdbc = db
	SetPoolC()
}

func SetPoolC() {
	m := global.Config.Mysql

	global.Mdbc.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	global.Mdbc.SetMaxOpenConns(m.MaxOpenConns)
	global.Mdbc.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}
