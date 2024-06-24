'use strict'  // this is strict mode in js

console.log(this) // in nonstrict - {}........ in strict mode - {}

function fn(){
    console.log(this) // in nonstrict - global object........ in strict mode - undefined
}
fn()

const ob = {
    name:"sai",
    fun: function(){
        console.log(this) //in nonstrict - ob itself........ in strict mode return ob again
    }
}
ob.fun()


const obj = {
    name:"sai",
    fun: function(){
           function fun2(){
            console.log(this) //in nonstrict - global object........ in strict mode return undefined
           }
           fun2()
    }
}
obj.fun()

// See the below pic to get this behaviour on browser too 

