// DataTypes

// 1. Primitive (value is passed)
//    7-types : String, Number, Boolean, null, undefined, Symbol, BigInt

//2. Reference Type or Non primitive
//   3-types: Array, Objects, Functions


// ******* the typeof operator ******8

console.log(typeof "sai") //string
console.log(typeof 23) //number
console.log(typeof false) //boolean
console.log(typeof null) //object (****anomaly****)
console.log(typeof undefined) //undefined
console.log(typeof 12345n) //bigint
console.log(typeof Symbol('123')) //symbol


const arr = [1,2,3]

const obj = {name:"sai"}

const fun =  () => {
    console.log("hii")
}

function fun2(){
    console.log("hii")
}

console.log(typeof arr) //object
console.log(typeof obj) //onject
console.log(typeof fun) //function
console.log(typeof fun2) //function