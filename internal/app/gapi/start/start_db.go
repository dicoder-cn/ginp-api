package start

import (
	"fmt"
	"ginp-api/configs"
	"ginp-api/internal/db/mysql"
)

func startDB() {
	mysql.InitDb(
		configs.MysqlIp(),
		configs.MysqlPort(),
		configs.MysqlUser(),
		configs.MysqlDb(),
		configs.MysqlPwd(),
	)

	//迁移表
	if mysql.GetWriteDb() != nil {
		//自动迁移表结构
		err := mysql.GetWriteDb().AutoMigrate(EntityAutoMigrateList...)
		if err != nil {
			fmt.Println("迁移表结构失败" + err.Error())
			panic(err)
		}
	}
}
