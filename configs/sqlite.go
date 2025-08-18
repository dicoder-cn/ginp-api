package configs

import "ginp-api/pkg/cfg"

// sqlite数据库文件路径
const ConfigKeySqliteDbPath = "sqlite.db_path"

// 设置默认值
func init() {
	cfg.SetDefault(ConfigKeySqliteDbPath, "data.db")
}

func SqliteDbPath() string {
	return cfg.GetString(ConfigKeySqliteDbPath)
}
