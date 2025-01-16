
# match case is like switch case and was introduced in 3.10 v python 

#no need to write break explicitly 
# no fallthrough clause also

day = 5

match day:
    case 1:
        print("Mon")
    case 2:
        print("Tue")
    case 3:
        print("Wed")
    case 4:
        print("Thu")
    case 5:
        print("Fri")
    case 6:
        print("Sat")
    case 7:
        print("Sun")
    case _:
        print("No proper day")



# or option 
day = 2

match day:
    case 1 | 2 | 3:
        print("Mon, Tue, or Wed")
    case 4:
        print("Thu")
    case 5:
        print("Fri")
    case 6:
        print("Sat")
    case 7:
        print("Sun")
    case _:
        print("No proper day")
