fact_list = []

def Pfactors(num):
    f = []
    for i in range(2,num+1):
        cnt = 0
        while num%i==0:
            num=num/i
            cnt+=1
        if cnt!=0:
            f.append([i,cnt])
    return f

for i in range(2,101):
    fact_list.append(Pfactors(i))
