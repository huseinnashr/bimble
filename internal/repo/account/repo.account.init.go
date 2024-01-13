package account

import (
	"github.com/huseinnashr/bimble/internal/config"
	"github.com/huseinnashr/bimble/internal/domain"
)

type Repo struct {
	config      *config.Config
	sqlDatabase domain.ISQLDatabase
	redis       domain.IRedis
}

func New(config *config.Config, sqlDatabase domain.ISQLDatabase, redis domain.IRedis) domain.IAccountRepo {
	return &Repo{
		config:      config,
		sqlDatabase: sqlDatabase,
		redis:       redis,
	}

}
