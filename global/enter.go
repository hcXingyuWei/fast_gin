package global

import (
	"fast_gin/configs"

	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const Version = "v1.0.0.1"

var (
	Config *configs.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Es     *elastic.Client
)
