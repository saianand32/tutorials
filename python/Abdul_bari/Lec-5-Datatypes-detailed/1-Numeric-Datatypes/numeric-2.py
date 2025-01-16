

#yes , bool and complex also are numeric

# 1. boolean - 

a = True
b = False

print(a) #True

print(type(a)) #class<'bool'>


print(int(a)) #1
print(int(b)) #0


#so in python 0 means False and any other number is Truthsy value

# #this goes in infinite loop
# while -9:
#     print(99)
# #this also goes in infinite loop
# while 3:
#     print(99)
#this doesnt run as 0 is False
while 0:
    print(99)


# 2. Complex ()in python we use j or J instaed of i

x = 2+3j
y = 2.33-5j
z = x+y

print(z) #(4.33-2j)

print(x.real) #2.0
print(x.imag) #3.0