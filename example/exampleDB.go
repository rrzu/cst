package example

import (
	"fmt"

	"github.com/rrzu/constMp"
)

// 展示单位
const (
	DBNone DB = iota
	DBMysql
	DBPostgres
	DBOracle
	DBHologres
	DBClickhouse
)

// DB 展示单位
type DB int

type MineDB *constMp.Cst[Database, any]

var db = &constMp.Cst[DB, constMp.Options[Database, any]]{
	Typ: constMp.DataTypeNumber,
	Words: constMp.Words[DB, constMp.Options[Database, any]]{
		{
			Value:  DBNone,
			CnName: "无",
		},
		{
			Value:  DBMysql,
			Group:  &constMp.Group{"OLTP"},
			CnName: "MySQL",
			Mine:   database.ToOptions(),
		},
		{
			Value:  DBPostgres,
			Group:  &constMp.Group{"OLTP"},
			CnName: "Postgres",
		},
		{
			Value:  DBOracle,
			Group:  &constMp.Group{"OLTP"},
			CnName: "Oracle",
		},
		{
			Value:  DBHologres,
			Group:  &constMp.Group{"OLAP"},
			CnName: "Hologres",
		},
		{
			Value:  DBClickhouse,
			Group:  &constMp.Group{"OLAP"},
			CnName: "Clickhouse",
		},
	},
}

func init() {
	fmt.Println("init db")
	constMp.Register(TypDB, db)
}
