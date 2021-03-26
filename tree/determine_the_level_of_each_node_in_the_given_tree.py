tree = dict()
tree[0]=[1,2]
tree[1]=[3,4]
tree[2]=[5,6]
tree[3]=[7,8]
tree[4]=[]
tree[5]=[]
tree[6]=[]
tree[7]=[]
tree[8]=[]
def determine_node_level_in_tree(tree,root):
    node_level=dict()
    queue = [root]
    node_level[root]=0
    while len(queue) != 0 :
        p = queue.pop(0)
        for n in tree[p] :
            queue.append(n)
            node_level[n]=node_level[p]+1
    return node_level
print(determine_node_level_in_tree(tree,0))