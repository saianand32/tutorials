

// 1. Defining generic functions

function identity<T>(val: T): T{
    return val
}


identity<number>(8);

// using interface type in the above method
interface Bottle{
    brand:string,
    type: number
}

let bottle1: Bottle = {brand:"xys", type:20}

identity<Bottle>(bottle1);


// 2. Arrow functions

const identity2 = <T>(val: T): T => {
    return val
}
