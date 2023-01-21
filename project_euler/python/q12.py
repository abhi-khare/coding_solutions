a = [0,1,2,2]
for i in range(4,10000):
    j = 2
    num = i
    temp = 1
    while j*j<=num:
        cnt = 1 
        while num%j==0:
            cnt = cnt + 1
            num = num/j
        temp = temp * cnt
        j = j+1
    if num!=1:
        temp*=2
    a.append(temp)

for i in range(0,len(a)):
    if a[i]*a[i+1]>500:
        print(i)
        break;  