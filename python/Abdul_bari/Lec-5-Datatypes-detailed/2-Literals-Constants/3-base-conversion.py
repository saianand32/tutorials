

a = 10


# 1. binary of a 
x = bin(a)
print(x)  #0b1010

x = f"{a:b}" #"1010" bin string

# 2. octal of a 
x = oct(a)
print(x) #0o12

x = f"{a:b}" #"12" bin string

# 3. hexadecimal of a 
x = hex(a)
print(x) #0xa

# 4. can also use bool type 
x = bin(True)
print(x) #ob1