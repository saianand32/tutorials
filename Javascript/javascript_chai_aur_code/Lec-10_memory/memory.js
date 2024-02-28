
// for  Primitive Stack is used 

// for reference/non primitive Heap used 


//copy is passed not a reference
let str = "sai"
let str2 = str
str2 = "anand" 
console.log(str) //sai
console.log()


//The reference is passed
let arr = [ 1,2,3,4]
arr2 = arr
arr2[2] = 99
console.log(arr)