// in JAVASCRIPT everything is either object or primitive
// ---Primitive - Number, string, boolean, null, undefined, symbol
// ---nonprimitive - objects(under objects- date, arrays, functions)

// We will demo how arrays and functions are objects too

const ob ={
    name:"sai",
    roll:32,
    age:21
}
for(let i in ob){ // use for in loop for object keys traversal
    console.log(i+":"+ob[i])
}

// **************************Arrays*********************************
const arr = [1,2,3,4,5];
console.log(arr)
arr.name = "array"
arr.para = "this is array"
console.log(arr) // this prints [ 1, 2, 3, 4, 5, name: 'array', para: 'this is array' ]

for(let i in arr){ 
    console.log(i+":"+arr[i])
} // if now i use a for in loop the 1,2,3,4,5 is mapped with respective indexes eg - 
// 0:1
// 1:2
// 2:3
// 3:4
// 4:5
// name:array
// para:this is array

const arr2 = [1,2,3,4];
arr2[40] = 99; // the element is inserted at index 40 and remaining middle ones are set to undefined
console.log(arr2) // it displays [ 1, 2, 3, 4, <36 empty items>, 99 ]
console.log(arr2[30])// it gives undefined


//************************** Functions ****************************************

// functions too behave like objects u can add key value pairs 

function fun(){
    console.log("fun is executed")
}

console.log(fun) //prints [Function: fun]

fun.college = "nmit"
fun.id = 39
console.log(fun) // prints [Function: fun] { college: 'nmit', id: 39 }

// can use a for in loop as well
for(let i in fun){ 
    console.log(i+":"+fun[i])
} 
//prints - only  the key value pairs but not the [Function: fun]
// college:nmit
// id:39

// the function is a speacial object with a separate code section in it which executes when called
 