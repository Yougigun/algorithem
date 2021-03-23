visited = set()
def depth_first_search(graph,node):
    if (node in visited): return 
    visited.add(node)
    print(node)
    for adjacent_node in graph[node]:
        depth_first_search(graph,adjacent_node)
        
graph = dict()
graph["4"] = ["50","7"]
graph["50"] = ["55","90"]
graph["7"] = ["87","2"]
graph["55"] = []
graph["90"] = []
graph["87"] = []
graph["2"] = []

depth_first_search(graph,"4")