//1. Defining Abstract classes (abstract classes also only in ts not in js)

abstract class A {
    private name: string // private protected allowed
    public age: number
    protected gender: string

    constructor(name: string, age: number, gender: string) {
        this.age = age
        this.name = name
        this.gender = gender
    }
    static strength: number //static allowed

    abstract fun(): string // creating an abstract function (same as java even if 1 abstract func then class must be declared abstract)


}

let ob1 = new A(); //************* err cannot create object of abstract class VVVIMP**********



class B extends A {
    //here unlike interfaces no need to redefine properties only abstract methods must be defined else err

    constructor( name: string, age: number, gender: string ) {
        super(name, age, gender);
    }

    fun() {
        console.log(A.strength) //access static variables 
        console.log(this.age) //correct
        console.log(this.name) //err as its private
        return "hii"
    }
}