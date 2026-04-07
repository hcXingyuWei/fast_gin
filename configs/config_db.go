package configs

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBMode string

const (
	DBMySQLMode  = "mysql"
	DBSqliteMode = "sqlite"
	DBPgsqlMode  = "pgsql"
)

type DB struct {
	Mode     DBMode `yaml:"mode"` //模式 mysql sqlite pgsql
	DBName   string `yaml:"db_name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (db DB) Dsn() gorm.Dialector {
	switch db.Mode {
	case DBMySQLMode:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.User,
			db.Password,
			db.Host,
			db.Port,
			db.DBName,
		)
		return mysql.Open(dsn)
	case DBSqliteMode:
		return sqlite.Open(db.DBName)
	case DBPgsqlMode:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			db.Host,
			db.User,
			db.Password,
			db.DBName,
			db.Port,
		)
		return postgres.Open(dsn)
	default:
		logrus.Warnf("未配置数据库连接")
		return nil
	}
	return nil
}
