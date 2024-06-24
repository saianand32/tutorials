

// 1. Synchronous callbacks *********************************************

function sayHii (name,callback){
  const msg = "Hii "+name;
  callback(msg);
}

function greet(msg){
console.log(msg)
}

sayHii('anand',greet)

sayHii('saishwar', (msg) => {
console.log(msg)
})


// 2. Asynchronous callbacks *******************************************

setTimeout(() => {
console.log("I am st 1")
}, 0)

setTimeout(() => {
console.log("I am st 2")
}, 0)
function sayBye(){
    console.log("bye")
}
 
sayBye() // here bye is printed then the timeouts are called
// it shows an async behaviour as the code does not wait for the timeouts