class BinHeap:
    def __init__(self, array):
        self.heapList = array
        self.currentSize = len(array)
        self.transformList()
        
    def push(self, node):
        """
        Insert a new item to heap
        (already heap.That,not unorder arr)
        """
        self.heapList.append(node)
        self.currentSize += 1
        index = self.currentSize - 1
        self.bubbleUp(index)
    def bubbleUp(self, i):
        parentIdx = abs(i - 1) // 2
        while i > 0 and self.heapList[i] < self.heapList[parentIdx]:
            self.heapList[i], self.heapList[parentIdx] = self.heapList[parentIdx], self.heapList[i]
            i = parentIdx
            parentIdx = (i - 1) // 2

    def pop(self):
        last = self.heapList[self.currentSize - 1]
        self.heapList[0] = last
        self.currentSize -= 1
        self.heapList.pop()
        self.bubbleDown()
    def bubbleDown(self, i=0):
        node = self.heapList[i]
        leftIdx = 2*i + 1
        rightIdx = 2*i + 2
        children = self.getChildren(leftIdx, rightIdx)
        smaller, smallerIdx = self.getSmaller(children)
        while children and node > smaller:
            self.heapList[i], self.heapList[smallerIdx] = self.heapList[smallerIdx], self.heapList[i]
            i = smallerIdx
            leftIdx = 2*i + 1
            rightIdx = 2*i + 2
            children = self.getChildren(leftIdx, rightIdx)
            smaller, smallerIdx = self.getSmaller(children)
    def getChildren(self, leftI, rightI):
        children = []
        if leftI < self.currentSize:
            children.append((self.heapList[leftI], leftI))
        if rightI < self.currentSize:
            children.append((self.heapList[rightI], rightI))
        return children
    def getSmaller(self, children):
        if len(children) == 0:
            return float('inf'), None
        if len(children) == 1:
            return children[0]
        else:
            childA, iA = children[0]
            childB, iB = children[1]
            return children[0] if childA < childB else children[1]
    
    def transformList(self):
        """
        start from  Last non-leaf node 
        (it'not as mentioned in article)
        """
        i = (self.currentSize - 1) // 2
        while (i >= 0):
            self.bubbleDown(i)
            i -= 1
arr=[1, 4, 6, 2, 5, 3, 9, 8, 7]
heap = BinHeap(arr)
print(heap.heapList)
print(id(heap.heapList),heap.heapList)
print(id(arr),arr)
