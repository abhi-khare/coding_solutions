sum = 0

def mult(l , x):
    carry = 0
    for i in range(0 , len(l)):
        temp = l[i]*2 + carry
        l[i] = temp % 10
        carry = int(temp / 10)
        
    return l
     

l = [0]*1000
l[0] = 1

for i in range(0,1000):
    l = mult(l,2)

for i in range(0,len(l)):
    sum = sum + l[i]

print(sum)