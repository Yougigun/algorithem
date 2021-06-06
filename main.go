package main

import (
	"fmt"
)

var a = map[string]struct{name string}{}
var b map[int]int

func main() {
	for i := range([]int{5,4,3,2,1}){

		fmt.Println(i)
	}
}
