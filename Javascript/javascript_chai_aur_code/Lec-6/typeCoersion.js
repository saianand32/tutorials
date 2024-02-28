
// 1. **** String to Number *****

let age = "20"
let ageNum = Number(age);
console.log(ageNum) //20
console.log(typeof ageNum) //number


// *** NaN ***
let height = "181cm" //even if the string is "abc" it will give NaN
let heightNum = Number(height)
console.log(heightNum) // NaN
console.log(typeof heightNum) //number

// Here see that typeof Nan is number 
console.log((NaN == NaN)) // false

// another way to convert strings to Numbers is using the unary + operator
let marks = "250"
let marksNum = + marks
console.log(marksNum)
console.log(typeof marksNum)

console.log(3+ +"5") //8

// 2. ***** undefined to Number   (it gives NaN) *****
let temp = undefined
let x = Number(temp)
console.log(x) //NaN
console.log(typeof x) //number

//3. ***** null to Number (it gives 0) *****
let temp2 = null
let x2 = Number(temp2)
console.log(x2) //0
console.log(typeof 2) //number

//4.***** Boolean to Number (converts to 0 or 1) *****
let flag = true
let flag2 = false
let flagNum = Number(flag)
let flagNum2 = Number(flag2)
console.log(flagNum+" "+flagNum2) // 1 0
console.log(typeof flagNum) //number

//5. Object or array to Number
let val11 = []
console.log(Number(val11)) //0
val11 = [0]
console.log(Number(val11)) // this also is 0
let val12 = {}
let val13 = {name:"sai"}
let val14 = [1,2,3]
console.log(Number(val12)) //NaN
console.log(Number(val13)) //NaN
console.log(Number(val14)) //NaN

// 6. ***** Number to Boolean ***** 
let val1 = 1
let val2 = 0
let val3 = -1
let val4 = 68
let val5 = -54
console.log(Boolean(val1)) //t
console.log(Boolean(val2)) //f
console.log(Boolean(val3)) //t
console.log(Boolean(val4)) //t
console.log(Boolean(val5)) //t
//So only 0 is false any other nonzero val is true even -ve elements are true

// 7. ***** String to Boolean ***** 
let v1 = ""
let v2 = "sai"
let v3 = "0"
console.log(Boolean(v1)) //false
console.log(Boolean(v2)) //true
console.log(Boolean(v3)) //true

//8. null to boolean
console.log(Boolean(null)) //false

//9. undefined to boolean
console.log(Boolean(undefined)) //false

//10. Object/arr to boolean
let v11 = [] //same with objects as arrays are basically objects
let v12 = [0]
let v13 = [1,2,3]
let v14 = {}
let v15 = {name:"sai"}
// console.log(Boolean(v11)) //t
// console.log(Boolean(v12)) //t
// console.log(Boolean(v13)) //t
// console.log(Boolean(v14)) //t
// console.log(Boolean(v15)) //t



// ********* Exceptions *********

// Easy way to remember exceptions 

// 1. Anything to Number 
// remember Number( [],[0],"","0",null ) => 0
// Number( [1,2,3],{},{name:"sai"},"abs","12g",undefined ) => NaN

//2. Anything to Boolean
// remember Boolean( "", 0, null, undefined, NaN) => false ..called falsey values 
// everything else true



// *** One More Left .... Anything to String ****
let value1 = 33
let value2 = []
let value3 = [1,3,4]
let value4 = {}
let value5 = {name:"sai"}
let value6 = undefined;
let value7 = null;

console.log(String(value1)) //"33"
console.log(String(value2)) //""
console.log(String(value3)) //"1,3,4"
console.log(String(value4)) //"[Object Object]"
console.log(String(value5)) //"[Object Object]"
console.log(String(value6)) //"undefined"
console.log(String(value7)) //"null"



