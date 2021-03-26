package bfs

type queue []int

func (q *queue) enqueue(v int) {
	*q = append(*q, v)
}

func (q *queue) dequeue() (v int) {
	v = (*q)[0]
	*q = (*q)[1:]
	return
}

func bfs(tree map[int][]int, root int) (res []int) {
	// init
	q := queue{}
	q.enqueue(root)
	for len(q) != 0 {
		p := q.dequeue()
		res = append(res, p)
		for _, v := range tree[p] {
			q.enqueue(v)
		}
	}
	return
}

func bfsNodeLevel(tree map[int][]int, root int) (res map[int]int) {
	res = map[int]int{root: 0}
	q := queue{root}
	for len(q) != 0 {
		p := q.dequeue()
		for _, n := range tree[p] {
			res[n] = res[p] + 1
			q.enqueue(n)
		}
	}
	return

}
