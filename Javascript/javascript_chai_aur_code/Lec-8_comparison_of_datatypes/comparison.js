// 1. Simple comparison 

console.log(1<2) //true
console.log(1>2) //false

console.log(1<=2) //true
console.log(1>=2) //false

console.log(2=="2") //true
console.log(2==="2") //false

//2. Comparison where datatypes not  same

// --- (a) 
console.log(2=="2") //true
console.log(2==="2") //false
console.log(2>"1") //true
console.log("2">1) //true
console.log("02">"9") //false


// --- (b)
//trick question
console.log(null > 0) //false
console.log(null == 0) //false
console.log(null >= 0) //true

console.log(undefined > 0) //false
console.log(undefined == 0) //false
console.log(undefined >= 0) //false

