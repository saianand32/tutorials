

// if u ant a property of a type to be readOnly

type User = {
    readonly _id: string, //readonly
    name: string,
    age: number,
    isActive: boolean,
    gender?: "string" //optional fields
}

let user1: User = {
    _id: "wejfbhiwebf7832y89",
    name: "sai",
    age: 7,
    isActive: true
}

user1.age = 90; // can change any non readonly properties
console.log(user1)

// but try to change _id

user1._id = "sihduwf" // error as u cannot change readOnly properties

export {}