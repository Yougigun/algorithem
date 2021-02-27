
def binarySearch(arr,tar):
    index = int(len(arr)/2)
    # print(index,arr)
    if  len(arr)==0: return False
    if arr[index]==tar:
        return True,index
    # r1 = binarySearch(arr[:index],tar)
    if binarySearch(arr[:index],tar):
        return True
    # r2 = binarySearch(arr[index+1:],tar)
    if binarySearch(arr[index+1:],tar):
        return True
    return False

arr = [1,2,3,4,5,100,5,4,3,6,7]
tar = 10
print(binarySearch(arr,tar))
    