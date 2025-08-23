package dbs

import (
	"ginpapi/internal/db/mysql"
	"ginpapi/internal/db/pgsql"
	"ginpapi/internal/db/sqlite"

	"gorm.io/gorm"
)

const (
	DbTypeMysql  = "mysql"
	DbTypePgsql  = "pgsql"
	DbTypeSqlite = "sqlite"
)

const useDbType = DbTypePgsql

var (
	DbRead  *gorm.DB
	DbWrite *gorm.DB
)

func GetReadDb() *gorm.DB {
	switch useDbType {
	case DbTypeMysql:
		return mysql.GetReadDb()
	case DbTypePgsql:
		return pgsql.GetReadDb()
	case DbTypeSqlite:
		db, err := sqlite.GetReadDb()
		if err != nil {
			panic(err)
		}
		return db
	default:
		return mysql.GetReadDb()
	}
}

func GetWriteDb() *gorm.DB {
	switch useDbType {
	case DbTypeMysql:
		return mysql.GetWriteDb()
	case DbTypePgsql:
		return pgsql.GetWriteDb()
	case DbTypeSqlite:
		db, err := sqlite.GetWriteDb()
		if err != nil {
			panic(err)
		}
		return db
	default:
		return mysql.GetWriteDb()
	}
}
