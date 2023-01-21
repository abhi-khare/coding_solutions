

sum = 1
a = 3
d = 2
for i in range(3,1002,2):
    sum += 4*a+6*d
    a += 4*d+2
    d +=2
    print(sum)

print(sum)
