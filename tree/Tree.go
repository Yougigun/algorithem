package main

import "fmt"

type Node struct {
	name string
	children []Node
}

type Tree struct{
	root Node
}

// Adjacency List
type TreeMap map[int][]int
var treeMap = TreeMap{
	0:[]int{1,2},
	1:[]int{3,4},
	2:[]int{5,6},
}
func main() {
	node1 := Node{"1",[]Node{}} 
	node2 := Node{"2",[]Node{node1}}
	tree := Tree{node2}
	fmt.Println(tree)
	fmt.Println(treeMap)



}