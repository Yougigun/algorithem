def bubble_sort(arr):
    n = len(arr)
    # find ith biggest
    for i in range(n-1,-1,-1):
        # only need to compare i-1 times
        for j in range(i):
            if arr[j]>arr[j+1]:
                temp = arr[j]
                arr[j]=arr[j+1]
                arr[j+1]=temp
            
    return arr


if __name__=="__main__":
    arr = [3,5,1]
    arr = bubble_sort(arr)
    print(arr)