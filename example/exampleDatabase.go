package example

import (
	"fmt"

	"github.com/rrzu/constMp"
)

const (
	DatabaseZulin   Database = "zulin"
	DatabaseServer  Database = "server"
	DatabaseProduct Database = "product"
)

type Database string

var database = &constMp.Cst[Database, any]{
	Typ: constMp.DataTypeString,
	Words: constMp.Words[Database, any]{
		{
			Value:  DatabaseZulin,
			CnName: "zulin",
		},
		{
			Value:  DatabaseServer,
			CnName: "server",
		},
		{
			Value:  DatabaseProduct,
			CnName: "product",
		},
	},
}

func init() {
	fmt.Println("init database")
	constMp.Register(TypDatabase, database)
}
