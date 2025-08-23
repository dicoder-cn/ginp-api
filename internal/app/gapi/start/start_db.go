package start

import (
	"fmt"
	"ginpapi/internal/db/dbs"
)

func startDB() {
	// 初始化PostgreSQL数据库
	dbs.InitDb(dbs.DbTypePgsql)

	//迁移表
	if dbs.GetWriteDb() != nil {
		//自动迁移表结构
		err := dbs.GetWriteDb().AutoMigrate(EntityAutoMigrateList...)
		if err != nil {
			fmt.Println("迁移表结构失败" + err.Error())
			panic(err)
		}
	}
}
