
out = [-1]*10000000

out[0] = 0
out[1] = 1
ans = 1
index = 1

def cal(i):
    
    if i>=1000000:
        if i%2==0:
            return cal(int(i/2))+1
        else:
            return cal(3*i+1) + 1 

    if out[i]!=-1:
        return out[i]
    else:
        if i%2==0:
            return cal(int(i/2)) + 1
        else:
            return cal(3*i+1) + 1


for i in range(2,1000000):
    out[i] = cal(i)
    if out[i]>ans:
        ans = out[i]
        index = i

print(index)