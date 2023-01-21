sieve = [1]*10000000

def sievef():
    sieve[0] = 0
    sieve[1] = 0
    for i in range(0,10000000):
        if sieve[i] == 1:
            for j in range(2*i,10000000,i):
                sieve[j] = 0

sievef()
for i in range(4,10000000):
    flag=1
    if sieve[i] == 0 and i%2==1:
        for j in range(3,i):
            if sieve[j]==1:
                num = int((i - j)/2)
                temp = int(pow(num,0.5))
                if temp*temp == num:
                    flag=0
                    break
        if flag ==1:
            print(i)
            break
