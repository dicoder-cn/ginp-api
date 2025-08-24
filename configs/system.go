package configs

import "ginp-api/pkg/cfg"

const (
	ConfigKeySystemAppName       = "system.app.name"
	ConfigKeySystemUserCenterUrl = "system.usercenter.url"
	//适用的数据库类型
	ConfigKeySystemDbType = "system.db.type"
)

const (
	defaultSystemAppName       = "dianji"
	defaultSystemUsercenterUrl = "http://localhost:8082"
	defaultSystemDbType        = "mysql" //pgsql,sqlite
)

func init() {
	cfg.SetDefault(ConfigKeySystemAppName, defaultSystemAppName)
	cfg.SetDefault(ConfigKeySystemUserCenterUrl, defaultSystemUsercenterUrl)
	cfg.SetDefault(ConfigKeySystemDbType, defaultSystemDbType)
}

func SystemAppName() string {
	return cfg.GetString(ConfigKeySystemAppName)
}

func SystemUserCenterUrl() string {
	return cfg.GetString(ConfigKeySystemUserCenterUrl)
}

func SystemDbType() string {
	return cfg.GetString(ConfigKeySystemDbType)
}
