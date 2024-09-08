package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.Get("mysql.user"),
		viper.Get("mysql.password"),
		viper.Get("mysql.host"),
		viper.Get("mysql.port"),
		viper.Get("mysql.db_name"),
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.max_open_conns"))
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.max_idle_conns"))

	return
}
func Close() {
	d, err := db.DB()
	if err != nil {
		zap.L().Error("close DB failed", zap.Error(err))
		return
	}
	d.Close()
}
