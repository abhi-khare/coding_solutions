def factor(num):
    l = []
    i = 2
    while(num>1):
        while(num%i==0):
            num=int(num/i)
            l.append(i)
        i +=1
    l = list(set(l))
    return len(l)

num = 2100
cnt = 1

while(True):
    num += 1
    if factor(num)>=4:
        cnt+=1
    else:
        cnt=0

    if cnt==4:
        print(num-3)
        break
