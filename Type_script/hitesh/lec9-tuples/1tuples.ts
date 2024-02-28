
// 1. simple array declaration

let arr: (string | number)[];

arr = [1,"sao",2]


// 2. tuples are more strict
// to make an order in the array

let tup: [number, string, boolean];

tup = [1,"sai", true]
// tup = [true, 1,"sai" ] //error as order matters


//3. Using type to define array and tuple

type UserArray = (string | number)[]


type UserTuple = [number, string]


//4. Anomalies
// tuples are not safe to push, unshift etc ops
// also u can reassign values and this is a problem

let tuple: UserTuple = [1,"sai"]

tuple[1] = "anand" // this is allowed and is wrong

// can push the defined in datatypes but not the others
// this also is a problem as tuple should be acually as how they are defined strictly
tuple.push(9)
tuple.push("anand")
// tuple.push(true) //error