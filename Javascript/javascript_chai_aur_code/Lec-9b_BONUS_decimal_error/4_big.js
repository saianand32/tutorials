const Big = require('big.js');


// In Big strict mode, creating a Big number from a primitive number is disallowed.

// Big.strict = true
// x = new Big(1)                         // TypeError: [big.js] Invalid number
// y = new Big('1.0000000000000001')
// y.toNumber()                           // Error: [big.js] Imprecise conversion

x = Big(0.6)
y = Big(0.7)
z = Big('1.000004536') //can also pass string
v = Big(1.899e-2) 

//1. plus
sum = x.plus(y).plus(z)
console.log(sum.toNumber()) //2.300004536

sum = Big(0.7).plus(0.8)
console.log(sum.toNumber()) //1.5

sum = Big(0.7).plus(v)
console.log(sum.toNumber()) //0.71899

//2. minus

sum = Big(0.7).minus(0.2)
console.log(sum.toNumber()) //0.5

// 3. times

prod = Big(0.7).times(0.2)
console.log(prod.toNumber()) //0.14

//4. div
div = Big(0.7).div(0.2)
console.log(div.toNumber()) //3.5

//5. mod
mod = Big(8.4).mod(2.4)
console.log(mod.toNumber()) //1.2

//6. pow (takes only whole numbers in pow() dont pass 0.5 etc)
pow = Big(9).pow(-5)
console.log(pow.toNumber()) //96.04

//7. sqrt

sqrt = Big(9).sqrt()
console.log(sqrt.toNumber()) //3

// MAJOR PROBLEM IS CANT DO DECIMAL POWERS



// 8. The value of a Big number is stored in a decimal floating point format in terms of a coefficient, exponent and sign.

x = new Big(-123.456);
x.c  // [1,2,3,4,5,6] coefficient (i.e. significand)
x.e // 2 exponent
x.s // -1 sign   



//9. Like JavaScript's Number type, there are toExponential, toFixed and toPrecision methods.

x = new Big(255.5)
console.log(x.toExponential(5)) // "2.55500e+2"
console.log(x.toFixed(5)) // "255.50000"
console.log(x.toPrecision(5)) // "255.50"


