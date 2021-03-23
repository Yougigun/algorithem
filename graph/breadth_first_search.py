visited = set()
candidate = [] # mock queue
def breadth_first_search(graph,node):
    candidate.append(node)
    while len(candidate)!=0:
        r = candidate.pop(0)
        visited.add(r)
        print(r)
        for n in graph[r]:
            if n not in visited:
                candidate.append(n)
                
graph = dict()
graph["4"] = ["50","7"]
graph["50"] = ["55","90"]
graph["7"] = ["87","2"]
graph["55"] = []
graph["90"] = []
graph["87"] = []
graph["2"] = []

breadth_first_search(graph,"4")