package bfs

import (
	"testing"
)

func TestBfs(t *testing.T) {
	//        0
	//    1       2
	//  3   4   5   6
	tree := map[int][]int{
		0: {1, 2},
		1: {3, 4},
		2: {5, 6},
		3: {},
		4: {},
		5: {},
		6: {},
	}
	ans := []int{0, 1, 2, 3, 4, 5, 6}
	size := len(ans)

	res := bfs(tree, 0)
	sizeRes := len(res)

	if size != sizeRes {
		t.Error(
			"For", "Size",
			"expected", size,
			"Got", sizeRes,
		)
	}

	for i, v := range ans {
		test := res[i]
		if test != v {
			t.Error(
				"For", "value",
				"expected", v,
				"Got", test,
			)
		}
	}
}

func TestNodeLevel(t *testing.T) {
	//        0
	//    1       2
	//  3   4
	tree := map[int][]int{
		0: {1, 2},
		1: {3, 4},
		2: {},
		3: {},
		4: {},
	}
	ans := []int{0: 0, 1: 1, 2: 1, 3: 2, 4: 2}
	res := bfsNodeLevel(tree, 0)
	size := len(ans)
	sizeRes := len(res)
	if size != sizeRes {
		t.Error(
			"For", "Size",
			"expected", size,
			"Got", sizeRes,
		)
	}
	for k, v := range ans {
		testV, ok := res[k]
		if !ok {
			t.Error("For node", k, "Expected", v, "Got", "nothing")
		}

		if v != testV {
			t.Error("For node", k, "Expected", v, "Got", testV)
		}
	}

}
