package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/anle/codebase/global"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)

	db, err := sql.Open("mysql", s)

	checkErrorPanic(err, "InitMysql initialization error")

	global.Logger.Info("Initializing MySQL Successfully")

	global.Mdb = db
	SetPool()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDB := global.Mdb

	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))

}
