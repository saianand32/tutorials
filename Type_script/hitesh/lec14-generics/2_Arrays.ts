
// 1. Generics with Arrays

interface Bottle{
    brand:string,
    type: number
}


function getProducts<T>(products: T[]): T{
    return products[0]
}




// 2. Another way to do it
const getProducts2 = <T>(products: Array<T>): T =>{
    return products[0]
}




// 3. using interface variable to call the function

const bottleArr: Bottle[] = [
    {brand:"xys", type:12},
    {brand:"xefw", type:9},
    {brand:"xywefs", type:72}
]

getProducts<Bottle>(bottleArr);



// 4. A common practice in React
// they add a comma <T,> to distinguish a component from a generic 

const getProducts3 = <T,>(products: Array<T>): T =>{
    return products[0]
}
