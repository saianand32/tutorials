const acc = 13432
let acc2 = 78
var ji = 90;

// var has a function scope
function fun(){
    var a = 900;
}
console.log(a) //error as a cant be accessed outside of  function

// let and const have block scopes 
{
    var a1 = "sai";
    let b = 900;
    const h = 8;
}
console.log(a1) //prints "sai"
console.log(b) //error as block scope
console.log(h) //error as block scope


// console.table() method
let name = "sai"
let accNum = 9875647
let age = 21
 let details = {
    name : name, 
    age : age, 
    accNum : accNum
}
console.table(details)

// U can also reserve memory for a variable without var, let const it behaves like var
city = "dehradun"
console.log(city) //prints dehradun (it also will have function scope)