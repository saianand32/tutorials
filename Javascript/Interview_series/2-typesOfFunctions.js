// 1. functions and return types and params***********************************************************************************************************************

const fun = (age) => {
console.log("age is "+ age)
}
fun(); // since I dint pass parameter it will print age is undefined

function fun2(age){
console.log("age is "+ age);
}
let ans = fun(22); // fun2 is called so it will print age is 22
console.log(ans) // here ans is undefined as no return val 



// 2. functions are treated as variables ************************************************************************************************************************
 // These are called function expressions
 
 const fnContainer = function fn(){
    console.log("hii")
 }

 fnContainer(); // prints hii
 // Note here i cannot call fn() independently as its assigned only can call fnContainer()
 // fn(); will give error not defined

 const fnContainer2 = function () {
    console.log("Im anonymous function")
 } // can totally skip naming the internal functions therefor it is called anonymous function



//  3. IIFE(Immediately Invoked Function Expression) *************************************************************************************************************
// NOTE - iife must be at the top of all functions in code else error
(function func () {
    console.log("Im iife function")
})();
// IIFE function is auto invoked when executed need not call below
console.log(988)


// 4. Arrow Functions (reduce syntax, react, this operator in binding) *******************************************************************************************

const arrfun1 = () => {
    return "yo dude";
}
const arrfun2 = () => console.log("hello world"); // single line arrowfunctions

const arrfun3 = () => "Hello world"; //Arrow Functions Return Value by Default

// Parameterized arrow functions
const arrfun4 = (price, units) => {
    return price*units;
}

const arrfun5  = (price, units) => price*units; // one line parameterized arrowfunc

const arrowfunc6 = age => age*5; // when only one parameter can skip the () brackets for the parameter

const arrowfunc7 = (x) => (x%2==0?true:false) // can skip return statement if use () or no brackets
const arrowfunc8 = (x) => { return x%2==0?true:false} // should write return if use {} else it will return undefined


// 5. Functions can be passed as parameters ***********************************************************************************************************************
// concepts of higher order functions, callbacks
const getPerson =() => console.log("saishwar") 

const sayHello = (paramfunc) => {
    let i = 90;
    console.log("hii"+paramfunc);
    console.log("hii",paramfunc);
    getPerson() // can be like a callback func
}

sayHello(getPerson);

// 5. Functions can be returned from functions ********************************************************************************************************************
// concept of closures

 function outer(){
   return function (){
    console.log("im inner")
   }
 }
 
 const temp = outer();
 temp();
 //OR*****
 outer()()



 const outerEs6 = () => {
    return ()=>console.log("im inner es6")
 }
 const temp2 = outerEs6();
 temp2();



