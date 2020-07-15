/**
 * @Author: duke
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2020/6/17 4:10 下午
 */

package mysql

import (
	"log"

	"lottery/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func LoadMysql(cfg setting.Config, lg *zap.Logger) {
	db, err := gorm.Open(cfg.Mysql.Source())
	if err != nil {
		lg.Error("Error connecting to mysql:", zap.Error(err))
		panic(err)
	}
	db.DB().SetMaxOpenConns(50)
	db.DB().SetMaxIdleConns(50)
	if cfg.Log.Level == "debug" {
		db.LogMode(true)
	}
	log.Println("Loading mysql is complete...")
}
