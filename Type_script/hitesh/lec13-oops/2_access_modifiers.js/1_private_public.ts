// 1. setting members as private/public 
// by default everything is public 

class User{
    public name: string
    private age: number = 9 //default value

    constructor(name: string, age: number){
        this.name = name
        this.age = age
    }

    getAge(){
        return this.age
    }
}

let sai = new User("sai", 20)

console.log(sai.age) //err as age is private
let saiAge = sai.getAge() //use a getter method to get age


//2. Another way to define private  use #<variable_name>

class User2{
    name: string
    #age: number

    constructor(name: string, age: number){
        this.name = name
        this.#age = age
    }

    getAge(){
        return this.#age
    }
}

let sai2 = new User2("sai", 20)

console.log(sai2.#age) //err as age is private 
let saiAge2 = sai2.getAge() //use a getter method to get age





// ************************* 3. Proffesional way to write classes in typescript **************************

class User3{

    constructor(public name: string, private age: number){} //note u need not do this.name = name & thi.age=age its auto done in ts if u use this kind of syntax

    getAge(){
        return this.age
    }
}

let sai3 = new User3("sai", 20)

let saiAge3 = sai3.getAge() //use a getter method to get age





export {}
