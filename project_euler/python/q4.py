
def isPerm(num):
    if num==int(str(num)[::-1]):
        return True
    return False

out=1;  
for i in range(100,1000):
    for j in range(i,1000):
        num=i*j;
        if isPerm(num):
            if out<num:
                out=num
        
print(out)