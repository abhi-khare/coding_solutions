x,i,max = 600851475143,2,1
while i*i<=x:
    while x%i==0:
        x=x/i;
        if i>max:
            max=i;
    i = i + 1
if x>max:
    max=x

print(max)
        
