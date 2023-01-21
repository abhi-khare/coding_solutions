
def hcf(a,b):
    if a==0:
        return b
    return hcf(b%a,a)

def lcm(a,b):
    return (a*b)/(hcf(a,b))


res=1

for i in range(1,21):
    res = lcm(res,i)

print(res)