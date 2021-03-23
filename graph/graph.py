class Graph:
    def __init__(self, nodes):
      self.nodes = nodes
class Node:
    def __init__(self, name,children):
      self.name = name
      self.children = children
      
      
graph = dict()
graph["1"] = ["2"]
graph["2"] = ["1","3","4"]
graph["3"] = ["1","2","4"]
graph["4"] = ["1","2"]