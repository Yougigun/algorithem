package dfs

import (
	"fmt"
	"testing"
)

func TestDfs(t *testing.T) {
	//    0
	//  1   2
	// 3 4
	//    5
	tree := map[int][]int{
		0: {1, 2},
		1: {3, 4},
		2: {},
		3: {},
		4: {5},
		5: {},
	}
	ans := []int{0, 1, 3, 4, 5, 2}
	sizeAns := len(ans)
	r := res{}
	r.dfs(tree, 0)
	sizeRes := len(r)
	if sizeAns != sizeRes {
		t.Error("For", "Size", "Expected", sizeAns, "Got", sizeRes)
	}

	for i, v := range ans {
		if testV := r[i]; v != testV {
			t.Error("For", "Value", "Expected", v, "Got", testV)
		}
	}

}

func TestFindPath(t *testing.T) {
	//    0
	//  1   2
	// 3 4
	//    5
	tree := map[int][]int{
		0: {1, 2},
		1: {3, 4},
		2: {},
		3: {},
		4: {5},
		5: {},
	}
	root:=0
	dest := 5
	ans := []int{0, 1, 4, 5}
	sizeAns := len(ans)
	p := path{}
	p.findPath(tree, root,dest)
	fmt.Println(p)
	sizePath := len(p)
	if sizeAns != sizePath {
		t.Error("For", "Size", "Expected", sizeAns, "Got", sizePath)
	}

	for i, v := range ans {
		if testV := p[i]; v != testV {
			t.Error("For", "Value", "Expected", v, "Got", testV)
		}
	}

}
