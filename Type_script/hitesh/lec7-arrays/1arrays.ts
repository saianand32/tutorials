
// 1. cannot define array at all like this 

let arr = []

arr.push("nnn") // err Argument of type 'string' is not assignable to parameter of type 'never'

// coz the arr is auto assigned to type never


//2. defining arr with typpe

let arr2 : string[] = []

arr2.push("sai")
arr2.push("anand")
arr2.push("nivia")

console.log(arr)

//3. if u want to push anything

let arr3: any [] = []

arr3.push(8)
arr3.push("sai")
arr3.push(true)


//4. Another way

let arr4 : Array<number> = []


// 5. using types with array

type Users = {
    name:string,
    age: number
}

let arr5: Users[] = [{name:"sai", age:20}]


// 6. defining 2d arrays 

let arr6: string[][] = [[],[]]


export {}

