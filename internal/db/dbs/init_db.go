package dbs

import (
	"ginpapi/configs"
	"ginpapi/internal/db/mysql"
	"ginpapi/internal/db/pgsql"
	"ginpapi/internal/db/sqlite"
)

func InitDb(dbType string) {
	switch dbType {
	case DbTypeMysql:
		initMysql()
	case DbTypePgsql:
		initPgsql()
	case DbTypeSqlite:
		initSqlite()
	default:
		panic("db type not support")
	}

}

func initMysql() {
	mysql.InitDb(
		configs.MysqlIp(),
		configs.MysqlPort(),
		configs.MysqlUser(),
		configs.MysqlDb(),
		configs.MysqlPwd(),
	)
}

func initPgsql() {
	pgsql.InitDb(
		configs.PgsqlIp(),
		configs.PgsqlPort(),
		configs.PgsqlUser(),
		configs.PgsqlDb(),
		configs.PgsqlPwd(),
	)
}

func initSqlite() {
	sqlite.InitdDb(configs.SqliteDbPath())
}
