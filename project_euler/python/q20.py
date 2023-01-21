fact = 1
sum = 0
for i in range(1,101):
    fact = fact*i

fact = str(fact)
for i in range(0,len(fact)):
    sum += int(fact[i])

print(sum)