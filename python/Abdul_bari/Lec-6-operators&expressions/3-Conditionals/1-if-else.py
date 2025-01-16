
# if elif and else 
# logical operators used in if conditions --- and, or, not

# Example variables
age = 20
has_ticket = True
has_vip_pass = False

# If-elif-else logic with 'and', 'or', and 'not'
if age < 18:
    print("Sorry, you are too young to enter.")
elif age >= 18 and (has_ticket or has_vip_pass):
    # Check if the person is 18 or older and has either a ticket or a VIP pass
    if not has_ticket and has_vip_pass:
        print("VIP access granted. Enjoy the event!")
    else:
        print("Regular access granted. Enjoy the event!")
else:
    print("You can't enter without a ticket or a VIP pass.")
