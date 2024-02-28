
// 1. Using extends keyword


function fun<T, U extends number>(val1: T, val2: U): object{
    return {
        val1,
        val2
    }
}

fun<string, string>("anand","sai"); //err as string doesnt extend number





// 2. more realistic example with interfaces

interface Database{
    con: string,
    username: string,
    password: string
}
interface Database2{
    con2: string,
    username2: string,
    password2: string
}

const fun2 = <T extends Database> (val: T): object => {
    return {}
}

// extending multiple interfaces
const fun3 = <T extends Database & Database2> (val: T): object => {
    return {}
}



// 3. Can do the same with type alias

type School = {
    schoolName: string
}

type User = {
    name: string;
    age: number;
} & School

type User2 = {
    name: string
}

const sai: User = {schoolName:"ssshss", name: "saishwar", age: 22}
const sai2 = {name: "saishwar"}


const fun4 = <T extends School>(val: T): object => {
    return {};
};

fun4<User>(sai);
fun4<User2>(sai2); //err as User2 doesnt have User interface as a part of it


// 4. Same with classes 

class MySchool{
    scName: string = "sssjj"
}


// this class can only take a type which extends MySchool class
class Student <T extends MySchool>{
    
}








export {}
