// 1. FUNCTION scope---------------------------------------------

function fun(){
    var a = 9;
    let b= 0;
    const c = 99;
    console.log(a+" "+b+" "+c)
}
fun();
// console.log(a+" "+b+" "+c) // here be it var, let or const they cant be accessed outside function this is a function scope
// this throws a reference error



// 2. LEXICAL scope----------------------------------------------

function parent(){
    var a = 10;
    let b= 11;
    const c = 12;

    function child(){
        console.log(a+" "+b+" "+c);
    }

    child(); // this will not throw reference error as a child will have access to all parent variables 
    // this is called a lexical scope of this child() function
}
parent();


// 3. CLOSURES ----------------------------------------------------

// A FUNCTION ALWAYS BUNDLED WITH ITS LEXICAL SCOPE FORMS A CLOSURE

function parent2(){
    var a = 100;
    let b= 110;
    const c = 120;

    function child2(){
        console.log(a+" "+b+" "+c);
    }

    return child2; // we are returning this internal child function
}

const catchChild2 = parent2();
catchChild2() 
// the shocking thing u will notice is that the parent has finished executing and is removed from the stack
// but the child still has access to its parents variables even after that 
// when catchChild2() is called it basically is executing only the child2() function.. parent2 is not even in the 
//picture.. but still all parent2 variables are accessed
// this is called a closeure 
//so ----- A FUNCTION ALWAYS BUNDLED WITH ITS LEXICAL SCOPE FORMS A CLOSURE


