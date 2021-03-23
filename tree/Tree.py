from typing import Dict, List


class Node:
    def __init__(self, name, children) -> None:
        self.name = name
        self.children = children
class Tree:
    def __init__(self,root) -> None:
        self.root = root
        
# Adjacency List
tree = dict()
tree["0"]=["1","2"]
tree["1"]=["3","4"]
tree["2"]=["5","6"]