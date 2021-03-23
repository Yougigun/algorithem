package main

import (
	"fmt"
)

var a = map[string]struct{name string}{}
var b map[int]int

func main() {
	a["1234"] = struct{name string}{"Gary"}
	s := a["1234"]
	c :=&s
	s.name="asd"
	fmt.Println(c, s)
}
