
//use const > let > var (try not using var at all coz of block and functional scope)


// 4 ways to declare variable 
const a = 9;
let b = 90;
var c = "sai"
d = "anand"

let fname; //auto assigns undefined


//1. var has function scope 

function fun(){
    var num1 = 90;
}
// console.log(num1) // reference err

if(2<9){
    var num2 = 90;
} 
console.log(num2) // var does not have block scope so accessible out of block as well 

// Note - Any variable within lexical scope of function is still accessible (eg below) not related to hoisting

var arr1 = []
let arr2 = []
const arr3 = []

function fun2(){
    console.log(arr1)
    console.log(arr2)
    console.log(arr3)
}

fun2()

//2.  let and const are block scope
{
    let aa = 90;
}
// console.log(aa) //ref err

// const has to be initialized when declared else same as let block scope

// 3. Functions scope of function 

function func1(){
    function func2(){
        console.log("fun2")
    }
}

func2() // cannot access






//******** HOISTING ******** Hoisting is a concept or behavior in JavaScript where the declaration of a function, variable, or class and imports goes to the top of the scope they were defined in

// 1. var hoisting 

console.log(d) //undefined as d is hoisted and initialized to undefined

var d ; // hoisted to top of function and initialized to undefined


// 2. let and const are hoisted to the top but not initialized to undefined so there will lie a TEMPERAL DEAD ZONE for them

// with let


console.log(tr)//temperal dead zone (ref err)
let t;

//const

// const g; //err need to init


// 3. Functions 
//functions are also hoisted like var but are not undefined like var .. they can be used







// declare and assign to multiplle variables same value

let n1, n2, n3
n1 = n2 = n3 = 90




