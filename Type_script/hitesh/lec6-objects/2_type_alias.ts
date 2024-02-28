
// 1. type aliases are like classes kind of

type User = {
    name: string,
    age: number,
    disabled: boolean
}


// now try the same anomaly

function createUser(user: User):User{
    return user
}

let user1 = createUser({name:"sai", age: 20, disabled: false}) //correct

let obj = {name:"sai", age: 20, disabled: false, salary: 200000}
let user2 = createUser({name:"sai", age: 20, disabled: false, salary: 200000}) //error as extra field there



// 2. can also rename primitive types  like this 

type bool = boolean
type num = number
type str = string

let myName:str = 2 // err as it should be string

export {}