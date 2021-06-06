def countsort(A):
    n = len(A)
    maxsize = max(A)
    
    # make a array to store the count with index as number
    carray = [0] * (maxsize + 1)
    for i in range(n):
        carray[A[i]] = carray[A[i]] + 1
        
    ## Harder to understang
    # i = 0
    # j = 0
    # while i < maxsize + 1:
    #     if carray[i] > 0:
    #         A[j] = i
    #         j = j + 1
    #         carray[i] = carray[i] - 1
    #     else:
    #         i = i + 1
    
    # Easy Readable
    pos = 0
    for i in range(maxsize):
        # carray[i] is the count of index.
        for _ in range(carray[i]):
            A[pos] = i
            pos+=1
            

# Time an Space complexity is O(max(Arr))
A = [3, 5, 8, 9, 6, 2, 3, 5, 5]
print('Original Array:',A)
countsort(A)
print('Sorted Array:',A)






