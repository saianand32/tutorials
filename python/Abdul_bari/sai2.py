


starting_amount = 180000

amount = starting_amount
age = 46
max_age = 76

lump_sum = 0

while age <= 76:
    print("Age: "+str(age)+", " + str(amount))
    lump_sum += amount*12
    amount = amount + ((5*amount)/100)
    age+=1

print("lumpSum: "+str(lump_sum))
