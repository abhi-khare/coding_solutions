arr = [1]*2000009
def sieve():
    arr[0]=0
    arr[1]=0
    for i in range(2,2000009):
        if arr[i]==1:
            j = 2*i
            while j<2000009:
                arr[j] = 0
                j = j + i
            

sum = 0 

sieve()
print(arr[0:30])
for i in range(0,2000000):
    if arr[i]==1:
        sum = sum + i

print(sum)