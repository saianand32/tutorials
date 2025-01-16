# Python program demonstrating break, continue, and pass in loops

# 1. Example using 'break'
def example_break():
    print("Example of 'break':")
    for i in range(10):
        if i == 5:
            print(f"Breaking out of the loop at i = {i}")
            break  # Exit the loop when i equals 5
        print(i)
    print("Loop ended\n")



# 2. Example using 'continue'
def example_continue():
    print("Example of 'continue':")
    for i in range(10):
        if i % 2 == 0:  # Skip even numbers
            continue  # Skip the rest of the loop for even numbers
        print(i)  # This will only print odd numbers
    print("Loop finished\n")



# 3. Example using 'pass'
def example_pass():
    print("Example of 'pass':")
    for i in range(5):
        if i == 3:
            pass  # Do nothing when i == 3
        else:
            print(f"Processing {i}")
    print("Loop finished\n")


# Main function to run the examples
def main():
    example_break()
    example_continue()
    example_pass()

# Run the main function
if __name__ == "__main__":
    main()
