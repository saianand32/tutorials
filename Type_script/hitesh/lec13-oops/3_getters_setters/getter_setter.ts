class User{
    public name: string
    private age: number

    constructor(name: string, age: number){
        this.name = name
        this.age = age
    }

    // this is a getter need to add get keyword in front and it must always return a val and must never have a parameter
    get getAge(): number{ 
        return this.age
    }

    // a setter cannot have any return type and only exactly compulsory one parameter
   public set setAge(newAge: number){
        this.age = newAge
    }
}


let sai = new User("sai", 22)

let age = sai.getAge() //err note a getter defined with get pretext cannot be called like a function 
age = sai.getAge //this is the correct way

sai.setAge(22) //err as this also is treated like properties
sai.setAge = 22 //this is the correct way






export {}