// even befor running the code js provides 3 major things - ****************************************************

// 1. Global object - (in Node - global object, in browser it is Window)
console.log(global) // in nodejs
console.log(window) // only in browser

// 2. This object
console.log(this) // in nodejs it is {} empty
console.log(this) // in browser it is equal to the window object

// 3. Code Section - 
// for code and data variables functions etc......



// ************************************ HOISTING OF VARIABLES **************************************

console.log("hii "+namee);
var namee;
console.log("yo "+namee);
namee = "saish";
console.log("hello "+namee)

// Basically the memory for var is allocated in creation phase -
// - In creation phase the execution context is set (global...etc), 'this' object is provided  
// and code section is created and memory allocated for var and assigned undefined if not assigned to a value;

// *************************************** HOISTING FUNCTIONS ***********************************
// NOTE ------ in JS function definitions are hoisted but not function declarations

// for functions the memory is allocated before the code runs therefore u can call function before declaration like
fun();

function fun () {
    console.log("im fun")
}

// Here memory is allocated in heap and a reference named fun is assigned to the memory location
// therefore we come accross another paradigm;
fun2();

function fun2 (){
    console.log("im first fun2")
}
function fun2 (){
    console.log("im second fun2")
}
function fun2 (){
    console.log("im third fun2")
}

fun2();

// here the memory is allocated in heap top to bottom so first a reference named fun2 is assigned to "im first fun2" one later 
// we come accross the same named reference allocated to another heap location of "im second fun2" so the reference is reassigned
// at end it finally is assigned to the latest definition in the code top-bottom


// ********************************* VAR , LET and CONST and TEMPORAL DEADZONE*****************************************
// 1. VAR 
// --- var is hoisted in js ->  
//     -- its declaration is hoisted and it is initialized to undefined untill the line of assignment like var s = "sai" ..etc

// 2. LET and CONST 
// --- let and const are also hoisted but -> 
//    -- for let and const declaration is hoisted but not assigned to anything not even undefined therefore they cause errors

// In JavaScript, let and const declarations are hoisted, but not their initialization.
// This means that when you use let or const to declare a variable, the declaration is hoisted to the top of the current block scope, 
// but the variable is not initialized until the line of code where the declaration appears. 
// This is often referred to as the "TEMPORAL DEADZONE".
    
   