package global

import (
	"database/sql"

	"github.com/anle/codebase/pkg/logger"
	"github.com/anle/codebase/setting"
	"github.com/dgraph-io/ristretto/v2"
	"github.com/redis/go-redis/v9"
)

var (
	Config     setting.Config
	Logger     *logger.LoggerZap
	Rdb        *redis.Client
	Mdb        *sql.DB
	LocalCache *ristretto.Cache[string, string]
)
