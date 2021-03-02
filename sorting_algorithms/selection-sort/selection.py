def selection_sort(arr):
    n = len(arr)
    result = []
    for _ in range(n-1):
        minVal = arr[0]
        minValIndex = 0
        for j,v in enumerate(arr):
            if v < minVal :
                minVal = v
                minValIndex = j
        arr.pop(minValIndex)
        result.append(minVal)
    result.append(arr[0])
        
                
                
    return result 
                
if __name__=="__main__":
    test_case = [3,2,1,1,5,3,4,6,7,0]
    arr = selection_sort(test_case)  
    print(arr)    
            
        
        