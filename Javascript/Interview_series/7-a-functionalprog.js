// in Functional programming We will see -------------
// 1. Imperative vs declarative way of coding
// 2. Pure vs Impure functions
// 3. Mutable vs Immutable code


// 1. Imperative vs Declarative*********************************

// --------Imperative way for odd even ------------------

const a = 4;
let b;
if(a%2 == 0) b=true;
else b = false;

console.log(b)

// -------Declarative way-----------------

const fun = (x) => (x%2==0?true:false)
console.log(fun(4))




// 2. Pure vs Impure functions (Pure function is a function that gives same results with same parameters) **************************88

// ----------impure functions------------------------

const v = 9;

function add(b){
    console.log(a+b); // clearly the function uses a variable from outer scope
}
// here the problem is that for the same parameter we can get different results by changing the outer variable, leads to ambiguity


// ----------Pure functions------------------------


function add(a,b){ // this is a pure function 
    console.log(a+b); // side-effect
}
// still the above function not properly implemented it has a side-effect
// side-effect means that the purefunction is manipulating something still out of scope of itself...
// here console.log is an outer aspect that is being run from the function we can remove side effect by--

function add(a,b){ // this is a pure function without side effects
    return (a+b); 
}
console.log(add(2,3))



// 3. Mutable and Immutabele

// ----Mutable-----------------------------------------

const ob1 = {
    name:"sai",
    roll:39
}
const ob2 = ob1;

console.log(ob1)
console.log(ob2) // both are same

ob2.name = "ankur" // this changes both ob1 and ob2 due to same reference
console.log(ob1)
console.log(ob2) 

// Thats why strings are immutable in java also cause of ambiguity

// Immutable ------------------------------------------

// can make the above immutable by using spread operator

const obj = {
    name:"sai",
    roll:39
}
// const obj2 = obj; // we wont do this
const obj2 = {...obj} // using the spread operator


obj2.name = "ankur" // this changes only ob2 
console.log(obj)
console.log(obj2) 


