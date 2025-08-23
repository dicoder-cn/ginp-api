package sdemotable

import (
	"ginpapi/internal/app/gapi/model/system/mdemotable"
	"ginpapi/internal/db/mysql"
)

var DemoTable *mdemotable.Model

func Model() *mdemotable.Model {
	if DemoTable == nil {
		dbRead := mysql.GetReadDb()
		dbWrite := mysql.GetWriteDb()
		DemoTable = mdemotable.NewModel(dbRead, dbWrite)
	}
	return DemoTable
}
