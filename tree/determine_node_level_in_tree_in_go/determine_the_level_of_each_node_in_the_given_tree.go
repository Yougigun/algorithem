package main

import "fmt"

func main() {
	test_tree := map[int][]int{
		0: {1, 2},
		1: {3, 4},
		2: {5, 6},
		3: {7, 8},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
	}
	fmt.Println(determine_node_level_in_tree(test_tree, 0))
}

func determine_node_level_in_tree(tree map[int][]int, root int) map[int]int {
	node_level := map[int]int{root: 0}
	queue := []int{}
	queue = append(queue, root) //enqueue
	var p int
	for len(queue) != 0 {
		p = queue[0]
		queue = queue[1:] //dequeue
		for _, n := range tree[p] {
			queue = append(queue, n)
			// key : decide level when adding candidate node in queue
			node_level[n] = node_level[p] + 1
		}
	}
	return node_level
}
