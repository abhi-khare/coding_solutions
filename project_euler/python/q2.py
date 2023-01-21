a=0
b=2
sum=2
while True:
    temp = 4*b + a
    if temp>4000000:
        break
    sum  = sum + temp
    a=b
    b=temp

print(sum)