// Create a class, constructor, functions 
class Person {
    constructor(name, age){
        this.name = name;
        this.age= age;
    }

    showDetails(){
        console.log("Hii " + this.name)
    }
}


class Child extends Person {
    constructor(name, age){ // can have only one constructor no constructor overloading in js
        super(name, age);
    }
    
}

let person1 = new Person('sai',21)
console.log(person1)
person1.showDetails()

let child = new Child('Keshav',32)
console.log(child)
child.showDetails()

// Inheritance is not suggested to be used in js

