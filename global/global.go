package global

import (
	"database/sql"

	"github.com/anle/codebase/pkg/logger"
	"github.com/anle/codebase/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

/*
Config
Redis
Mysql
...
*/

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
	Mdbc   *sql.DB
)
