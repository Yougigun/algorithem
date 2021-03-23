class Node:
    def __init__(self, name, left, right):
        self.name = name
        self.left = left
        self.right = right


class BinaryTree:
    def __init__(self, root):
        self.root = root


def make_complete_tree(nodes):
    initNodes = [Node(name, None, None) for name in nodes]
    for i in range(len(nodes)):
        node = initNodes[i]
        if (2*i+1) <= (len(nodes)-1):
            node.left = initNodes[2*i+1]
        if (2*i+2) <= (len(nodes)-1):
            node.right = initNodes[2*i+2]
    tree = BinaryTree(initNodes[0])
    return tree


def make_complete_tree_map(nodes):
    tree = dict()
    size = len(nodes)
    for i in range(size):
        tree[nodes[i]] = [None, None]
        if (2*i+1) <= (size-1):
            tree[nodes[i]][0] = nodes[2*i+1]
        if (2*i+2) <= (size-1):
            tree[nodes[i]][1] = nodes[2*i+2]
    return tree


def in_order_traversal(node):
    if node != None:
        in_order_traversal(node.left)
        print(node.name)
        in_order_traversal(node.right)


def in_order_traversal_tree_map(node, tree):
    # print(node)
    if node != None:
        in_order_traversal_tree_map(tree[node][0], tree)
        print(node)
        in_order_traversal_tree_map(tree[node][1], tree)


def pre_order_traversal(node):
    if node != None:
        print(node.name)
        pre_order_traversal(node.left)
        pre_order_traversal(node.right)

def pre_order_traversal_tree_map(node, tree):
    # print(node)
    if node != None:
        print(node)
        pre_order_traversal_tree_map(tree[node][0], tree)
        pre_order_traversal_tree_map(tree[node][1], tree)

def post_order_traversal(node):
    if node != None:
        post_order_traversal(node.left)
        post_order_traversal(node.right)
        print(node.name)

def post_order_traversal_tree_map(node, tree):
    # print(node)
    if node != None:
        post_order_traversal_tree_map(tree[node][0], tree)
        post_order_traversal_tree_map(tree[node][1], tree)
        print(node)

if __name__ == "__main__":
    nodes = ["10", "5", "20", "9", "18", "3","7"]
    
    # in-order traversal 
    print("in-order traversal")
    tree = make_complete_tree(nodes)
    in_order_traversal(tree.root)
    print("map-tree")
    tree = make_complete_tree_map(nodes)
    root = "10"
    in_order_traversal_tree_map(root, tree)
    
    # pre-order traversal
    print("pre-order traversal")
    tree = make_complete_tree(nodes)
    pre_order_traversal(tree.root)
    print("map-tree")
    tree = make_complete_tree_map(nodes)
    root = "10"
    pre_order_traversal_tree_map(root, tree)
    
    # pre-order traversal
    print("post-order traversal")
    tree = make_complete_tree(nodes)
    post_order_traversal(tree.root)
    print("map-tree")
    tree = make_complete_tree_map(nodes)
    root = "10"
    post_order_traversal_tree_map(root, tree)
    
    node1 = Node("11",None,None)
    node2 = Node("11",None,None)
    print(node2 is node1)