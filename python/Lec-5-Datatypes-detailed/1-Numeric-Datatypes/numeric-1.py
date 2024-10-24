


# 1. int

# in python there is no size limit for numbers can be as large as possible
# as everything is an object in python so handling memory is different from other languages 

# size of a number is variable as per length 

x = 123
print(x.__sizeof__()) #28

y = 1337376192387394629374623984729384723094234
print(y.__sizeof__()) #44



# 2. float 
# same like int this also infinite size 

x = 123.33
print(x.__sizeof__()) #24


# scientific representation 
y = 12.333e10 
print(y) #123330000000.0
print(y.__sizeof__()) #24

z = 12.333e-10
print(z) #1.2333e-09


#formatted printing 
x = 3.55444
print(f"{x:.3f}") #3.554



