const Decimal = require('decimal.js');

//Out of the 3 use this and only this

x = Decimal(0.6)
y = Decimal(0.7)
z = Decimal('1.000004536') //can also pass string
v = Decimal(1.899e-2) 

//1. plus
sum = x.plus(y).plus(z)
console.log(sum.toNumber()) //2.300004536

sum = Decimal(0.7).plus(0.8)
console.log(sum.toNumber()) //1.5

sum = Decimal(0.7).plus(v)
console.log(sum.toNumber()) //0.71899

//2. minus

sum = Decimal(0.7).minus(0.2)
console.log(sum.toNumber()) //0.5

// 3. times

prod = Decimal(0.7).times(0.2)
console.log(prod.toNumber()) //0.14

//4. div
div = Decimal(0.7).div(0.2)
console.log(div.toNumber()) //3.5

//5. mod
mod = Decimal(8.4).mod(2.4)
console.log(mod.toNumber()) //1.2

//6. pow (can pass any fraction or whole number)
pow = Decimal(9).pow(0.5)
console.log(pow.toNumber()) //3
pow = Decimal(9).pow(-5)
console.log(pow.toNumber()) //96.04

//7. sqrt

sqrt = Decimal(9).sqrt()
console.log(sqrt.toNumber()) //3


//8. Like JavaScript's Number type, there are toExponential, toFixed and toPrecision methods.
x = new Decimal(255.5)
console.log(x.toExponential(5)) // "2.55500e+2"
console.log(x.toFixed(5)) // "255.50000"
console.log(x.toPrecision(5)) // "255.50"


//9 ceil and floor
x = new Decimal(255.5)
console.log(x.ceil()) //256
console.log(x.floor()) //255


//10 can also check out toBinary, octal etc also cosine, sine functions on 
//https://www.npmjs.com/package/decimal.js/v/10.4.3
