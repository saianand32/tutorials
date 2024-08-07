// types: number, string, boolean, null, undefined, object, array, symbol, tuple,void, never, unknown, any

// sytax let var_name: type = value



// 1. number 

let age: number = 99;
let salary: number = 25000.56;

//2. string
let greetings: string = "saishwar"

// greetings = true // will give error as greetings is a type string 


// 3. Boolean 
let flag: boolean = true;
 

//4. any keyword behaves as a dynamic type
let sai: any = 90
sai=true;
sai="anand"


let uidd: unknown;

// The any and unknown types in TypeScript serve somewhat similar purposes but have different behaviors and use cases:

// any:

// The any type effectively opts out of type checking in TypeScript. It can hold any value, and TypeScript will not provide type-checking or inference on operations involving any.
// Using any effectively disables TypeScript's type system for the values it is applied to.
// It's commonly used when interfacing with JavaScript libraries or when the type of a variable cannot be known in advance.
// It's less type-safe compared to unknown because TypeScript will allow any operations on variables of type any without type checking.

// unknown:

// The unknown type is more restrictive than any. It represents a value about which the type-checker knows nothing or little. Unlike any, you cannot perform arbitrary operations on values of type unknown.
// Values of type unknown can be assigned to any other type without type errors, but you must perform type-checks or type assertions before operating on them.
// It's often used in scenarios where you want to enforce stricter type checking but still need flexibility in handling values of uncertain types, such as when consuming external data or dealing with dynamic input.
// Using unknown forces you to perform type-checks or type assertions to safely operate on values, which leads to better type safety compared to any.





//** typescript type inference (auto)

let uid = 123;
console.log(uid)
// uid=false; //gives false as it auto assigns number type when uid is declared and assigned 

//** auto type inference for any

let hero;

function getHero() {
    return true;
}

hero = getHero()

//here the variable hero is implicitly assigned any as ts cant determine the type of it



// 5. null, undefined

let name: undefined;
let fname: null;


//6. bigint

let sai32: bigint = 232334234n;



export {}

