// 1.

var s1 = "saishwar";

const fun = () => {
  console.log(s1);
};

fun();

// 2.

var s2 = "anand";
const fun2 = () => {  
  console.log(s2);
  var s2 = "utsav";
  console.log(s2)
};
fun2();

// In the above 2 examples we see that whenever a function is called a new execution context is created for the function
// so when fun2 is called a new var sai is hoisted to top and so the first is undefined then "utsav" is printed

// whereas in the first example since no var s1 is there in the execution context of fun() therefore the var s1 is seeked from
// the parent execution context...its just like oops overriding kindof

// ****************************************** SCOPE AND LEXICAL SCOPE ********************************************

// 1. The scope of a variable in JavaScript determines where the variable can be accessed in the code. 
// In JavaScript, there are two types of variable scopes: - global and local

// 2. What do u mean by lexical scope
console.log("line 1", varName);
var varName = 10;

function fun_b(){
console.log("line 2", varName)
}


console.log("line 3", varName);

function fun_a(){

    console.log("line 4", varName)
    var varName = 20;

    fun_b();

    console.log("line 5", varName);
}

fun_a();

// a-- In the above example initially concider main fn scope so var undefined hoisted on top therefore "line 1 undefined" is printed
// b-- then we come to "line 3 10" here var will be 10 
// c-- Now fun_a() is called and its own execution context is created with varname=undefined hoisted therefore "line 4 undefined "printed
// d-- now varName is assigned 20 but in scope of fun_a() and then we call fun_b()

// ******* now very important ****************
// e-- Lexical scope means if variable not exists in the function then it looks above the function definition for it.
//     here in fun_b() varname is not there so it takes varname = 10 from main fn scope which is just above it
//     also remember that lexical scope is seen with respect to the function definition and not function call
//     here thats why we get varName = 10 and not varName = 20;
//     if say fun_b() was defined within fun_a() and called then lexically above it would be fun_a()'s varname=20 but here 
//     the lexically above fn is the main one or global execution context so varname = 10;

// ******************* Scope Chain **************************
// - the process of looking for the variable in lexically above scope till u reach the global execution context scope 
// is termed as a scope chain in js (see image 5-b-scopechain.png)

// *************************** VARIABLE SHADOWING *******************
// when a same named variable in a block shadows one in outer block 
// 1. legal variable shadowing
let sa=99;
{
let sa=88
console.log(sa) //88 is printed sa=99 is shadowed
}
console.log(sa)


// 2. illegal variable shadowing
// this is illegal as var here has function scope therefore js treats it like its trying to redeclare a let
let ss=99;
{
var ss=88
console.log(ss) 
}
console.log(ss)

