
cnt = 13
prime=[1]*1000000
def shifter(num):
    temp = num[1:]
    temp = temp + num[0]
    return temp

def sieve():
    prime[0] = 0
    prime[1] = 0
    for i in range(2,1000000):
        if prime[i] == 1:
            for j in range(2*i,1000000,i):
                prime[j] = 0

sieve()
for i in range(100,999999):
    num = str(i)
    flag = 1
    for j in range(0,len(num)):
        num = (shifter(num))
        x = int(num) 
        if x%2==0:
            flag=0
            break;
        else:
            if prime[x]!=1:
                flag=0
                break
    if flag:
        cnt += 1
print(cnt)
