// 1. Record utility 

type keyValType = Record<number, string>


let obj: keyValType = {
    1: "sai",
    2: "anand",
    "jj": 9 //errr as key value pair do not matc the type
}

//the above can be used to define an object type with a particular type of keyValue pairs

let obj2: Record<number, boolean> = {
    1: false,
    4: true
}
