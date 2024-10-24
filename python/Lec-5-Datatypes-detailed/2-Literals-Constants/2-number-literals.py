

# For integer literals in number system 

a = 10 #decimal
b = 0b1010 #binary
c = 0B1010 #binary
d = 0o12 #octal
e = 0O12 #octal
f=0xA #hexadecimal
g=0XA #hexadecimal

print(a,b,c,d,e,f,g) #10 10 10 10 10 10 10


# h = 0b11.0b11 #error (not allowed)

# (Above all is same in Golang)
 
# in complex number only real part u can write in any number system but imag part must be decimal number system
i = 0b11 + 7j

#Reading from keyboard in another number system
price = int(input("\nenter price"), 2) #enter input 0b101
print(price) #5