// 1. in  js

let ob = {
    name: "sai",
    age: 89
}

console.log("age" in ob) //prints true


//2. in ts u can use it with interface variables annd object type aliases

interface User{
    name: string
}

interface Admin{
    isAdmin: boolean
}

function isAdmin(account: User | Admin){
    if('isAdmin' in account){ //narrowing the type
        return true
    }
    return false
}

 
// can do with types also (but works on object types only)

type User2 = {
    name: string
}

type Admin2 = {
    isAdmin: boolean
}

function isAdmin2(account: User2 | Admin2){
    if('isAdmin' in account){ //narrowing the type
        return true
    }
    return false
}