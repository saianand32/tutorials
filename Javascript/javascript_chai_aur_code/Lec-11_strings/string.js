//Declare String

let str = "saishwar"

let str2 = new String("anand")

// 2. String interpollation (prefer this over concat)
let age = 12
let str3 = `my age is ${age}`

//3. string functions

let myStr = "saishwar anand"

//1. length
console.log(myStr.length) //14

//2. upper case and lower case ... to capitalize use logic  no function
//these return a value n dont change the actual string
console.log(myStr.toUpperCase()) // SAISHWAR ANAND
console.log(myStr.toLowerCase()) // saishwar anand

//3. charAt and indexof 
console.log(myStr.charAt(2)) //i
console.log(myStr.indexOf('r')) //7
console.log(myStr.lastIndexOf('an')) //11

//4. substring
let newStr = myStr.substring(0,3)
console.log(newStr) //sai

newStr = myStr.substring(3)
console.log(newStr) //shwar anand

//5. slice (like substring but can take -ve values as well)
// remember the negative index works but still the first param should be smaller index than the 2nd
//its  never printing the string in reverse from -2 to 0 it  will give empty string

newStr = myStr.slice(0,-3) 
console.log(newStr) //saishwar an

newStr = myStr.slice(-10, 8) 
console.log(newStr) //hwar


//6. reverse (not inbuilt for string)
myStr = "saishwar anand"
revStr = myStr.split('').reverse().join('')
console.log(revStr) //dnana rawhsias


//7. trim(removes start & end spaces & \n)
myStr = "   sai anand   "
trimmedStr = myStr.trim()
console.log(myStr)     //   sai anand
console.log(trimmedStr)//sai anand

startTrimmedStr = myStr.trimStart() 
endTrimmedStr = myStr.trimEnd()


//8. replace

let url = "https://google.com/auth%20/login"
let url2 = url.replace('%20', '-')
 console.log(url)//https://google.com/auth%20/login
console.log(url2)//https://google.com/auth-/login


//9. includes (returns boolean)
myStr = "sai anand"
console.log(myStr.includes('ai')) //true
console.log(myStr.includes('ais')) //false


//10. split syntax split(separator, limit)
myStr = "saishwar anand"
revStr = myStr.split('',4).reverse().join('')
console.log(revStr) //sias


//11. charCodeAt to get ascii val
myStr="saish"
let asciiNum = myStr.charCodeAt(2)
console.log(asciiNum) //105

let char = String.fromCharCode(asciiNum);
console.log(char); //i
