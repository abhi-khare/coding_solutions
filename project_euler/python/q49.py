primes = [1]*10000
Plist = []
def sieve():
    primes[0],primes[1] = 0,0
    for i in range(2,10000):
        if primes[i]==1:
            if i>999:
                Plist.append(i)
            for j in range(2*i,10000,i):
                primes[j]=0

def isPerm(x,y,z):
    x = sorted(str(x))
    y = sorted(str(y))
    z = sorted(str(z))
    for i in range(0,4):
        if x[i]==y[i] and y[i] == z[i]:
            continue
        else:
            return False
    return True

sieve()
for i in range(0,len(Plist)):
    for j in range(i+1,len(Plist)):
        mid = int((Plist[i]+Plist[j])/2)
        if primes[mid]==1 and isPerm(Plist[i],Plist[j],mid):
            print(Plist[i],mid,Plist[j])
