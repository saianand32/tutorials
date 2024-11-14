
# In Python, the else block attached to a for or while loop is executed only if the loop completes normally—i.e., 
# without encountering a break statement. If the loop is terminated by a break, the else block is skipped.


#the else here will execute if and only if break never gets called... if loop breaks else wont execute

i = 1

while i < 10:
    print(i)
    i+=1
    if i > 5:
        break

else:
    print('all 10 nums got printed')

print('end of program') 

# output  (else didnt execute cause loop used the break clause when i == 6)
# 1
# 2
# 3
# 4
# 5
# end of program

i = 1
while i < 10:
    print(i)
    i+=1
    if i > 15:
        break

else:
    print('all 10 nums got printed')

print('end of program') 

# output  (else executed cause loop never used the break clause)
# 1
# 2
# 3
# 4
# 5
# 6
# 7
# 8
# 9
# all 10 nums got printed
# end of program