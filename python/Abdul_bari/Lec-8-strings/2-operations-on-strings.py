

# 1. concat

# a - using + 
s1 = "sai"
s2 = "anand"
s3 = s1+s2
print(s3) # saianand

# b  - using " " space
s3 = "hello" "world"
print(s3) # helloworld

# but this wont wprk 
# s3 = s1 s2  #error the space one will work with raw strings not variables

# s3 = "hello"+15 #error as cannot  use + on str and int
s3 = "hello"+ str(15)



# 2. Repetition 
s1 = "sai"
s2 = s1*10
print(s2) # saisaisaisaisaisaisaisaisaisai


# 3. indexing 
# 0,1,2,3,4,5   -1,-2,-3,-4,-5,-6 

# 4. Slicing 
s1 = "hello world"

# [start:end:step]

s2 = s1[0:4:2]
print(s2) #hl

# 5. Reverse 
s1 = "saiswar"

s2 = s1[::-1]
print(s2) # rawsias

s2 = "".join(list(reversed(list(s1)))) #using reversed method (dont use )
print(s2) # rawsias

# 5. in and not in 

s1 = "saishwar anand"
print("sai" in s1) #True
print("hello" in s1) #False