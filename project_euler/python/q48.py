

out = 0

def modexp(x,y,mod):
    if x==1 or y==0:
        return 1
    elif y==1:
        return x%mod
    else:
        z = modexp(x,int(y/2),mod)
        if y%2==1:
            return ((z%mod)*(z%mod)*(y%mod))%mod
        else:
            return ((z%mod)*(z%mod))%mod


for i in range(1,1001):
    num = modexp(i,i,10000000000)
    #num = pow(i,i,10000000000)
    out = (out + num)%10000000000

print(out)
