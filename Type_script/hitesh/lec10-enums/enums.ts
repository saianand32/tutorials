// Enum in Ts

enum seatChoice {
    AISLE, //0
    MIDDLE, //1
    WINDOW //2
}

const mySeat = seatChoice.WINDOW

// note the first in enum gets assigned val = 0... u can redefine the val

enum seatChoice2 {
    AISLE = 'aisle',
    MIDDLE = 'middle',
    WINDOW = 'window'
}

// the above produces a lot of js code


// use const to produce less
const enum seatChoice3 {
    AISLE = 'aisle',
    MIDDLE = 'middle',
    WINDOW = 'window'
}