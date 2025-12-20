package example

import "github.com/rrzu/constMp"

const (
	TypCPActivityStatus constMp.Typ[CPActivityStatus, any]              = "CPActivityStatus"
	TypBasicIndicator   constMp.Typ[BasicIndicator, MineBasicIndicator] = "BasicIndicator"
	TypShowUnit         constMp.Typ[ShowUnit, any]                      = "ShowUnit"
	TypDB               constMp.Typ[DB, constMp.Options[Database, any]] = "DB"
	TypDatabase         constMp.Typ[Database, any]                      = "Database"
)
