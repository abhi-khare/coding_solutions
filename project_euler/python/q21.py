l = [-1]*10001

l[0] = 0 
l[1] = 0
sum = 0

def Pdivsum(x):
    i = 2
    out = 1
    while i*i<=x:
        if x%i ==0 and int(i)!=int(x/i):
            out += int(i) + int(x/i)
        elif x%i==0 and int(i) == int(x/i):
            out += int(i)
        i+=1
    
    return out

for i in range(2,10001):
    if l[i] != -1:
        continue
    temp = Pdivsum(i)
    if temp <= i:
        continue
    else:
        temp2 = Pdivsum(temp)
        if temp <=10000:
            l[i] = temp
        if temp <=10000:
            l[temp] = temp2 
        if i == temp2:
            sum += i + temp
print(sum)

