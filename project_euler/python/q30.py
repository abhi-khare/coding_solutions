


res = 0

for i in range(10,999999):
    num = str(i)
    sum = 0
    for j in range(0,len(num)):
        sum += pow(int(num[j]),5)
    if i ==sum:
        res+=i

print(res)
