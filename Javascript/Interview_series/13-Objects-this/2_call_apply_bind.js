// CALL APPLY AND BIND *********************************************************

// use of call apply bind is so that we can use one function for many objects no new function declarations for objects
// 1. CALL ---------------------------------------------

const obj1 = {
    name:"saishwar",
    roll:39
}

const obj2 = {
    name:"ankur",
    roll:22
}


function fun(){
    console.log(this.name+" "+this.roll)
}

function fun2(city, car){
    console.log(this.name+" "+this.roll+" "+city+" "+car)
}

fun.call(obj1);                       // the call method can be used to call a function for various objects
fun2.call(obj2, "Bangalore","BMW" )


// 2. APPLY--------------------------------------------------------

fun2.apply(obj2,['Bangalore',"ferari"]) // the only difference between call and apply is u have to pass arguments as an array in apply method.

// 3. BIND----------------------------------------------------------

const boundFunction = fun2.bind(obj1, "Dehradun","Mustang" );
boundFunction()
// the bind method stores the function in another and u can call that later 

