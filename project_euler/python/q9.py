
for i in range(1,1000):
    for j in range(i+1,1001):
        if 1000000 + 2*i*j - 2000*(i+j)==0:
            print(i*j*(1000-i-j))
