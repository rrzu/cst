package main

import (
	"encoding/json"
	"fmt"

	"github.com/rrzu/constMp"
	"github.com/rrzu/cst/example"
)

func main() {
	var a []byte

	// a, _ = json.Marshal(constMp.Get(example.TypCPActivityStatus).ToOptions())
	// fmt.Println(string(a))
	// a, _ = json.Marshal(constMp.Get(example.TypBasicIndicator).ToOptions())
	// fmt.Println(string(a))
	// a, _ = json.Marshal(constMp.Get(example.TypShowUnit).ToOptions())
	// fmt.Println(string(a))
	//
	// a, _ = json.Marshal(constMp.Get(example.TypDB))
	// fmt.Println(string(a))

	fmt.Println("here")
	a, _ = json.Marshal(constMp.Get(example.TypDB))
	fmt.Println("1:", string(a))

	// a, _ = json.Marshal(constMp.GetFilterGroup(example.TypDB, "OLAP"))
	// fmt.Println("2:", string(a))
	//
	// a, _ = json.Marshal(constMp.Get(example.TypDB))
	// fmt.Println("3:", string(a))
	//
	// a, _ = json.Marshal(constMp.GetFilterGroup(example.TypDB, "OLTP"))
	// fmt.Println("4:", string(a))
}
