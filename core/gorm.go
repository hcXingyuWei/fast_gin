package core

import (
	"fast_gin/global"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitGorm() (db *gorm.DB) {
	cfg := global.Config.DB
	dialector := cfg.Dsn()
	if dialector == nil {
		return
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //不生成实体外键
	})
	if err != nil {
		logrus.Fatalf("%s数据库连接失败：%s", cfg.Mode, err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		logrus.Fatalf("获取%s数据库失败:%s", cfg.Mode, err)
	}
	//设置连接池
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(30 * time.Second)
	sqlDb.SetMaxOpenConns(100)

	logrus.Infof("%s数据库连接成功", cfg.Mode)
	return
}
