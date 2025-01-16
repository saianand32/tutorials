

s1 = "alaska"
s2 = "canada"

# when u compare it gets compared in dictionary order 
# also caps get more preference than small

if s1 < s2:
    print("s1<s2") #this prints
else:
    print("s2<s1")

s1 = "alaska"
s2 = "aLaska"
if s1 < s2:
    print("s1<s2")
else:
    print("s2<s1") #this prints

# 2. Never compare number strings 
s1 = "21"
s2="121"
print(s1<s2) #False 
