package dfs

type res []int
func (r *res)dfs(tree map[int][]int,p int){
	*r = append(*r , p)
	for _ , n := range tree[p] {
		r.dfs(tree,n)
	}
	return
}
type path []int
func (p *path)findPath(tree map[int][]int,node , dest int) (ok bool){
	*p = append(*p , node)
	if (node==dest) {return true}
	for _ , n := range tree[node] {
		ok = p.findPath(tree,n,dest)
		if ok { return true}
		*p = (*p)[:len(*p)-1] //delete when go one up layer 
	}
	return 
}