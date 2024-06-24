// 2 ways to demonstrate currying *******************************************************

//Currying is a technique in functional programming that performs the transformation
//of a function with multiple arguments into several functions containing a single argument in a sequence

// 1. Using Bind method ------------------------------------------

function sum(a, b, c) {
    console.log(a + b + c)
}

// const newSumFunctionWithDefaultParameters = sum.bind(sum,4);
// OR
const newSumFunctionWithDefaultParameters = sum.bind(this, 4); // here in sum now we take a = 4 as default parameter

newSumFunctionWithDefaultParameters(2, 3); // prints 9



//2. Using Closures-------------------------------------------------

function prodOuter(a) {

    function prodInner(b) {
        console.log(a * b)
    }
    return prodInner
}

const newProductFunction = prodOuter(3); // here we give the first element a default of 3

newProductFunction(2) // Prints 6...as it remembers the variables in its lexical environment ft closure

