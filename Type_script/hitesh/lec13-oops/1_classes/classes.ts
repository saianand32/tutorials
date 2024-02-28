class User {
    readonly email: string
    name: string

    constructor(email: string, name: string, age: number) {
        this.email = email
        this.name = name
        this.age = age //err as you didnt define age variable as class variable
    }

}

const sai = new User("saianand@gmail.com", "sai", 20);
sai.email = "sai@gmal.com" //err as readonly


export {}