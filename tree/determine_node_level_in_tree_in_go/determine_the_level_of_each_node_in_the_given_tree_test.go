package main

import "testing"

func TestNodeLevel(t *testing.T){
	test_tree := map[int][]int{
		0: {1, 2},
		1: {3, 4},
		2: {},
		3: {},
		4: {},
	}

	answer := map[int]int{
		0:0,
		1:1,
		2:1,
		3:2,
		4:2,
	}
	testResult := determine_node_level_in_tree(test_tree,0)
	
	size := len(answer)
	testSize :=len(testResult)

	if  testSize != size {
		t.Error(
			"For:","Size",
			"Expected:",size,
			"got:",testSize,
		)
	}

	for k, v := range testResult {
		if (answer[k] != v){
			t.Error(
				"For:",k,
				"Expected:",answer[k],
				"Got:",v,
			)
		} 
	}
}

func BenchmarkNodeLevel(b *testing.B){
	test_tree := map[int][]int{
		0: {1, 2},
		1: {3, 4},
		2: {},
		3: {},
		4: {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		determine_node_level_in_tree(test_tree,0)
	}
}