 "use strict"; // treat all JS code as newer version of js

// 1. Number  
let maxPossibleVal = Number.MAX_VALUE; //(1.79E+308)
let maxPossibleSafeIntValue = Number.MAX_SAFE_INTEGER; // (2^53 -1)

let minPossibleVal = Number.MIN_VALUE; //(5E-324)
let minPossibleSafeIntValue = Number.MIN_SAFE_INTEGER;  // (-1)*(2^53 -1)

//2. String

//3. Boolean

// 4. null  (standalone value ie. missing value etc..)

// 5. undefined (value not yet assigned)

// 6. BigInt 
let n = 934859472n;
let n2 = BigInt(123453)

// 7. Symbol (uniqueness)
 const id = Symbol('123')
 const id2 = Symbol('123')
 console.log(id1==id2) //false as Symbol makes uniqueness

// Non primitive 
// 1. Object , Array, Functions 

// The typeof operator*********
console.log(typeof 9) //number
console.log(typeof undefined) //undefined
console.log(typeof null) //object