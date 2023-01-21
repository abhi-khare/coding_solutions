



n = 40
k = 20

def nChoseK(n,k):
    if k == 0:
        return 1
    else:
        return (n*nChoseK(n-1,k-1))/k;


res = nChoseK(n,k)
print(res)