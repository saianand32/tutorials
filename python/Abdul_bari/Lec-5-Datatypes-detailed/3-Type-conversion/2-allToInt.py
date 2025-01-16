

a = 12.24
b = "123"
b1="ssdw3" #not possible to convert to int
c = True
d = 2+9j #not possible to convert to int

x = int(a)
print(x) #12

x = int(b)
print(x) #123

x = int(c)
print(x) #1


e = "0b101"
x = int(e, 2) #provide base radix
print(x) #5

