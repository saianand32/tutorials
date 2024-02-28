// protected is accessible in inherited classes like a private variable


class School{
    public schoolName: string = "ssshss"
    protected schoolAge: number
    private schoolRegion: string
    static population: number = 20000 //writing static variables

    constructor(schoolName: string,  schoolAge: number, region: string ){
        this.schoolAge =  schoolAge
        this.schoolName = schoolName
        this.schoolRegion = region
    }

}

class Student extends School {
    name: string
    gender: string

    constructor(name: string, gender: string){
        super("ssshss",80,"south")
        this.name = name
        this.gender = gender
    }

    setSchoolData(): void{
        this.schoolName = "lll" //can be done
        this.schoolAge = 90 //can be done as protected is accessible in the inherited class
        this.schoolRegion = "jbw" //err cant be done as private not accessible
    }
}

let sai = new Student("saish", "male")

console.log(sai.schoolName) // public so i can access and is inherited
console.log(sai.schoolAge)    //err protected is inherited but also i cant directly access need to use methods
console.log(sai.schoolRegion) //err private cant access as not inherited


//accesssing static variables u can use any
console.log(Student.population) // 20000
console.log(School.population) //20000