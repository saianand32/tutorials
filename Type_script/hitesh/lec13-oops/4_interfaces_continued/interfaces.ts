// Note interface only exists in typescript not in js


// 1. Common Mistakes in interfaces and access modifiers 
interface TakePhoto {
    cameraMode: string 
    filter: string
    burst: number = 90 //err as u cannot intialize variables in interface 
    //also no default methods in interfaces

    static fun(): void //err cannot have static/private/protected members (be it variables/methods) in interfaces
    private fun2(): void //err cannot have static/private/protected members (be it variables/methods) in interfaces
    protected fun3(): void //err cannot have static/private/protected members (be it variables/methods) in interfaces
    public fun4(): void // this also err as u cannot use any access specifiers in interface

    //..... basically cant use access specifiers or static in interface...
    //  by default everything is public but never us the public keyword
}


// implementing the interface 
class Instagram implements TakePhoto{ //err
    
}

//err above as js is not like java
// in java interfaces have all properties(variables)as public static final by default .. in js not like this
// also in java the class that implements interface must only define methods in the class the properties can be accessed by this.property_name or InterfaceName.property_name
// but in js both the properties and methods need to be defined in the class thats why error




// 2. Correct way 

interface TakePhoto2 {
    cameraMode: string 
    filter: string
    burst: number

    fun(): string
}

interface Story{
    createStory(): string
}


class Instagram2 implements TakePhoto2, Story{ 
    public cameraMode: string
    public filter: string
           burst: number

    fun(){
        return "Success"
    }

    public createStory(): string { 
        return "story created"
    }

    // Note: as u can see above the properties and methods of the implemented interfaces can be given either public or no modifier 
    //(cannot give private/protected as scope will reduce similar to java)
}
