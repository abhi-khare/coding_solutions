


res = 0

fact = [1,1,2,6,24,120,720,5040,40320,362880]


for i in range(10,9999999):
    num = str(i)
    sum = 0
    for j in range(0,len(num)):
        sum += fact[int(num[j])]
    if i == sum:
        res+=i

print(res)
