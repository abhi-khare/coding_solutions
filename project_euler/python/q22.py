
text = open('p022_names.txt','r').read()
text = text.split(',')
alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
char2index = {}
data=[]
i,out = 0,0

for names in text:
    data.append(names.replace('"',''))

for i in range(1,27):
    char2index[alphabet[i-1]] = i

data.sort()
print(data)

for name in data:
    sum = 0
    for alphabet in name:
        sum += char2index[alphabet]
    out += sum * i
    i += 1

print(out)