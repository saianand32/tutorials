// inAccuracy is there with decimal math in js  due to how numbers are stored

console.log(0.6+0.7) //1.2999999999999998
console.log(0.6+.7)  //1.2999999999999998

console.log(100.90 - 20.30) // 80.60000000000001
console.log(2232.00 * 0.1) //223.20000000000002


// How to avoid these 

// 1. Convert to non decimal and then convert back 
// very tedious process 


//2. Use Libraries (see next file)
