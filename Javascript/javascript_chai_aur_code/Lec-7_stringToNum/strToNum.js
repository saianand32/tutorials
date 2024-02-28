
//Why str to num is confusing



// 1. string concat with numbers 
console.log("1"+2) //12
console.log("2"+1) //21
console.log(1+2+"3") //33


 // 2. Convert any to Number using unary '+' operator

 console.log(+true) //1
 console.log(+false) //0
 console.log(+"") //0
 console.log(+"20") //20
 console.log(+"sai") //NaN
 console.log(+[]) //0
 console.log(+[1,2]) //NaN
 console.log(+{}) //NaN
 console.log(+{name:"sai"}) //NaN