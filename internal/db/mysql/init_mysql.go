package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb(ip, port, userName, dbName, dbPwd string) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", userName, dbPwd, ip, port, dbName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, //日志参数
	})

	if err != nil {
		fmt.Println("mysql链接失败" + err.Error())
		panic(err)
	}

}

func GetReadDb() *gorm.DB {
	//返回数据库实例的副本
	copyDb := *db
	return &copyDb
}

// GetDbInstance 获取gorm示例的副本
func GetWriteDb() *gorm.DB {
	//返回数据库实例的副本
	copyDb := *db
	return &copyDb
}
