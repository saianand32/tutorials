
//Note the  typeinference is not easy in functions so 
//always write the types :type in functions... 
//can skip this in variables declaration

// 1. type on params

function addNums (num1: number, num2: number){
    return num1+num2
}

console.log(addNums(2,4))

const addNums2 = (num1: number, num2: number) => {
    return num1+num2
}


console.log(addNums2(2,4))


// 2. type on return type

function addNums3 (num1: number, num2: number) : number {
    return num1+num2
}


console.log(addNums3(2,4))

const addNums4 = (num1: number, num2: number) : number => {
    return num1+num2
    // return num1+num2+"" //err as reurning string 
}


console.log(addNums4(2,4))


// 3. Default Params

function signupUser (id: number, age: number, dept: string = "ise"){
    console.log(id+age+dept)
}

signupUser(2,21)
signupUser(2,21,"ece")


// 4. using types with map function
let arr = [1,2,3,4]

arr.map((ele: number): string => {
    return ele.toString()
})


//5. void functions

const fun = (num: number): void => {
    console.log(num)
}

//6. return type never
// A function returning 'never' cannot have a reachable end point.

const errHandle = (errMsg: string): never => {
    throw new Error(errMsg)
}



//7. function with array

function funn(id: number, arr: string []) : number[]{
    return [1,2,3]
}

//8. function wih object

function funn2(id: number, obj :{name:string, roll: number}): {name: string, age: number}{
    return {name:"sai", age:21}
}



export {}