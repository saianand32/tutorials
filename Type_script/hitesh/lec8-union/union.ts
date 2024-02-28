
// 1. union type

let score: number | string = 8;
score = "jii"
// score = true  //error



// 2. union with types

type user = {
    name:string,
    age: number
}

type admin = {
    adminName: string
}

let sai : user | admin = {name:"sai", age: 20}


// 3. union with functions 

const getData = (id: number | string): boolean | string => {
    return true;
}


// 4. Anomaly 

const getDataa   = (id: number | string) => {
    // say u use .toUpperCase but it can be both number or string so error
    // let temp = id.toUpperCase()

    //need to do 
    if(typeof id === "string"){
        let temp = id.toLocaleUpperCase() // no error
    }

}


// 5. union with arrays  vvvimp 

// --- (i) array of either all strings or all numbers 
const data: string[] | number [] = [1,2, 9]
const data2: string[] | number [] = ["sai", "anand"]
const data3: string[] | number [] = ["sai", 2,"anand"] //error as its mixed

// --- (ii) array of either strings or numbers any combo (MOST IMPORTANT)
const data4: (string | number) [] = [1,2, 9, "sai", "anand"]

// --- (iii) commmon mistake
let data5: string | number[] ; // here it can be either string type "" or "number array" number[]



// 6. assigning constants 

let pi: 3.14 = 3.14;
pi = 9; //error

let seat: "aisle" | "middle" | "window";
seat = "window"