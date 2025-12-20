package example

import (
	"fmt"

	"github.com/rrzu/constMp"
)

// 展示单位
const (
	ShowUnitNone      ShowUnit = ""
	ShowUnitPercent   ShowUnit = "%"
	ShowUnitMoneyYUAN ShowUnit = "元"
	ShowUnitTimeH     ShowUnit = "h"
)

// ShowUnit 展示单位
type ShowUnit string

var showUnit = &constMp.Cst[ShowUnit, any]{
	Typ: constMp.DataTypeString,
	Words: constMp.Words[ShowUnit, any]{
		{
			Value:  ShowUnitNone,
			CnName: "无",
		},
		{
			Value:  ShowUnitPercent,
			CnName: "%",
		},
		{
			Value:  ShowUnitMoneyYUAN,
			CnName: "元",
		},
		{
			Value:  ShowUnitTimeH,
			CnName: "h",
		},
	},
}

func init() {
	fmt.Println("init showUnit")
	constMp.Register(TypShowUnit, showUnit)
}
