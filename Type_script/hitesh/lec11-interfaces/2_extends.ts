// 1. extending other interfaces 
interface Student{
    name: string,
    age: number
}

interface School{
    schoolName: string
}


interface labStudent extends School, Student{
    labName: string
}


let labStudent1: labStudent = {
    name:"sai",
    age: 20,
    schoolName: "xbnss",
    labName:"chem lab"
}