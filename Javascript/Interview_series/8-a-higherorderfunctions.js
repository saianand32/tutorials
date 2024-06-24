// ************************* HIGHER ORDER Functions*****************************

// A “higher-order function” is a function that accepts functions as 
// parameters and/or returns a function.

// 1. Map Function ------------------------
const arr = [1,2,3,4,5]

const mappedArray = arr.map((element, index) => {
    return element*element;
})

console.log(mappedArray)

// 2. Filter function --------------------------
const filteredArray = arr.filter((element, index) => {
    return element%2==0;
})

const filteredArrayy = arr.filter((element, index) => {
    if(element%2==0) return element;
})

console.log(filteredArray)


// 3. Reduce -----------------------------------

const sumArray = arr.reduce((accumulator,element, index) => {
return accumulator+element;
},0) ;// here we initialize the accumulator to 0

console.log(sumArray)

