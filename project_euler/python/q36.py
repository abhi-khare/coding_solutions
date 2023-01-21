sum = 0

def isPalindrome(num):
    num = str(num)
    i,j = 0,len(num)-1
    while i<=j:
        if num[i]!=num[j]:
            return False
        i+=1
        j-=1
    return True

def baseito2(num):
    return bin(num)[2:]


for i in range(1,1000000):
    if isPalindrome(i) and isPalindrome(baseito2(i)):
        sum += i

print(sum)
