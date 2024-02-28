// 1. Creating object type 

const user: object = {}

//this is any kind of object like [], function or object{} with any number of key vals
// to specify better read below



// 2. Creating objects with functions 


function createUser(obj: {name:string, age:number}){
    return obj
}


//can skip the obj: as well but dont use this
function createUserrrr({name:string, age:number}){
    
}

let user1 = createUser({name:"sai", age:20})
let user2 = createUser({name:"sai", age:19, roll:20}) //error coz roll is not supposed to be there


// but there is a flaw
let temp = {name:"sai", age:19, roll:20}
let user3 = createUser(temp) // but here no error even though roll is there

createUserrrr({name:"sai", age:89})



// 3. create object with type (not adviced due to same anomaly use type aliases)


let userr: {name:string, age: number} = {name:"sai", age:90}


// same anomaly
let userr2 : {name:string, age: number} ;

userr2 = {name:"sai", age:90, flag: true} //error on flag as it doesnt exist in type definition

//but if i create another object and assign userr2 to that obj this check doesnt work
let ob = {name:"sai", age:90, flag: true}
userr2 = ob; // err is gone but type check is not done so anomaly

export {}