arr = [0]*28123
num = []
sum = 0

def is_abundant(num):
    out = 1
    i = 2
    while i*i<=num:
        if num%i==0:
            out += i + num/i
        i +=1
    if (i-1)*(i-1)==num:
        out -= num
    if out>num:
        return True
    return False

for i in range(12,28123):
    if is_abundant(i):
        num.append(i)

for i in num:
    for j in num:
        if i+j <28123:
            arr[i+j] = 1

for i in range(0,28123):
    if arr[i]==0:
        sum += i

print(sum)