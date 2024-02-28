//1. interface definition
interface User{
    email: string,
    userId: number,
    readonly dbId: number,
    googleId?: string //optional field
}

const sai: User = {email:"", userId: 89, dbId:9900}



//2. Defining functions in interface (same as in types)
interface Student{
    name: string,
    getStudent : () => string //1st way
    getMarks() : number //2nd way

    getDetails: (id: number) => string [] //parametrized function way 1
    getDetails2(id: number) : string [] //parametrized function way 2
    
}



export {}