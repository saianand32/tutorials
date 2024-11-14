# Example 1: Basic while loop
print("Basic while loop:")
count = 0
while count < 5:
    print(count)
    count += 1  # Increment the counter
print("\n")

# Example 2: Infinite while loop (with break)
print("Infinite while loop with break:")
while True:
    user_input = input("Type 'exit' to quit: ")
    if user_input == 'exit':
        print("Exiting the loop...")
        break  # Exit the loop when 'exit' is typed
print("\n")

# Example 3: while loop with else
print("while loop with else:")
count = 0
while count < 3:
    print(count)
    count += 1
else:
    print("The loop completed without a break.")
print("\n")

# Example 4: Basic for loop (iterating over a list)
print("Basic for loop (iterating over a list):")
fruits = ["apple", "banana", "cherry"]
for fruit in fruits:
    print(fruit)
print("\n")

# Example 5: for loop with range() (iterating over a range of numbers)
print("for loop with range():")
for i in range(5):  # Loops from 0 to 4
    print(i)
print("\n")

# Example 6: for loop with range(start, stop, step) (custom range)
print("for loop with range(start, stop, step):")
for i in range(1, 10, 2):  # Loops from 1 to 9 with a step of 2
    print(i)
print("\n")

# Example 7: for loop with else
print("for loop with else:")
for i in range(5):
    print(i)
else:
    print("The loop completed all iterations.")
print("\n")

# Example 8: for loop with break (breaks the loop early)
print("for loop with break:")
for i in range(5):
    if i == 3:
        print("Breaking the loop at i =", i)
        break  # Exit the loop when i is 3
    print(i)
print("\n")

# Example 9: for loop with continue (skips an iteration)
print("for loop with continue:")
for i in range(5):
    if i == 2:
        continue  # Skip the rest of the loop when i is 2
    print(i)
print("\n")

# Example 10: for loop with enumerate() (get index and value)
print("for loop with enumerate():")
fruits = ["apple", "banana", "cherry"]
for index, fruit in enumerate(fruits):
    print(f"Index: {index}, Fruit: {fruit}")
print("\n")

# Example 11: for loop in reverse (using reversed())
print("for loop in reverse (using reversed()):")
fruits = ["apple", "banana", "cherry"]
for fruit in reversed(fruits):
    print(fruit)
print("\n")

# Example 12: for loop in reverse (using range with negative step)
print("for loop in reverse (using range with negative step):")
for i in range(4, -1, -1):  # Loops from 4 down to 0
    print(fruits[i])
print("\n")

