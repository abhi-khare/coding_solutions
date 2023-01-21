import math
s = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
alphaDict = {}
cnt = 0
for i in range(0,len(s)):
    alphaDict[s[i]]=i+1
print(alphaDict)

def is_square(integer):
    root = math.sqrt(integer)
    if int(root + 0.5) ** 2 == integer:
        return True
    else:
        return False

word_list = open('p042_words.txt').read().split(',')
print(word_list)
for word in word_list:
    val = 0
    for char in word:
        if char == '"':
            continue
        else:
            val += alphaDict[char]
    if is_square(1+8*val)==1:
        cnt +=1
print(cnt)
