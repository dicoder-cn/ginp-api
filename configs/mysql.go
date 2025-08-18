package configs

import "ginp-api/pkg/cfg"

const ConfigKeyMysqlPort = "mysql.port"
const ConfigKeyMysqlIp = "mysql.ip"
const ConfigKeyMysqlUser = "mysql.user"
const ConfigKeyMysqlPwd = "mysql.pwd"
const ConfigKeyMysqlDb = "mysql.dbname"

// 设置默认值
func init() {
	cfg.SetDefault(ConfigKeyMysqlIp, "127.0.0.1")
	cfg.SetDefault(ConfigKeyMysqlPort, "3306")
	cfg.SetDefault(ConfigKeyMysqlUser, "root")
	cfg.SetDefault(ConfigKeyMysqlDb, "")
	cfg.SetDefault(ConfigKeyMysqlPwd, "123456")
}

func MysqlIp() string {
	return cfg.GetString(ConfigKeyMysqlIp)
}
func MysqlPort() string {
	return cfg.GetString(ConfigKeyMysqlPort)
}

func MysqlUser() string {
	return cfg.GetString(ConfigKeyMysqlUser)
}
func MysqlPwd() string {
	return cfg.GetString(ConfigKeyMysqlPwd)
}
func MysqlDb() string {
	return cfg.GetString(ConfigKeyMysqlDb)
}
