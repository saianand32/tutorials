const fs = require('fs')

// 1. Parallel tasks with callbacks ******************************************

fs.readFile("f1.txt", (err, data) => {
    console.log(" f1data is - "+data)
})

fs.readFile('f2.txt', (err, data) => {
    console.log("f2 data is - "+data)
})

fs.readFile('f3.txt', (err, data) => {
    console.log("f3 data is - "+data)
})

// the above code seems to execute random order therefore its parallel



// 2. Serial tasks with callbacks***********************************************
  // to demo this ill write the callbacks separately

  function callback1 (err, data) {
    console.log("f1 data is - "+data)

    fs.readFile('f2.txt', callback2)
  }

  function callback2 (err,data){
    console.log("f2 data is - "+data)
    fs.readFile('f3.txt', callback3)
  }

  function callback3 (err,data){
    console.log("f3 data is - "+data)
  }

  fs.readFile('f1.txt', callback1) 

  // THis is nothing but a callback hell as nesting of callbacks can be seen its a flaw and to overcome this
  // we will look into promises
