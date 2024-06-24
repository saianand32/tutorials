// Objects 

let ob = {
    name:'sai',
    age: 20
}

console.log(ob['name'])
console.log(ob['age'])
console.log(ob.name)



// delete object properties

delete ob['age']
delete ob.age
console.log(ob)


//method vs function
// a method is a function within object 

ob = {
    name: "sai",
    getAge: function (){
        return 21
    }
}

console.log(ob.getAge())


//this keyword

ob = {
    fname: "sai",
    lname: "anand",
    age: 21,
    getSummary: function(){
        // return `${this.fname}, ${this.lname}, ${this.age}`
        return this //returns entire obj
    } 
}

console.log(ob.getSummary()) /// prints ob

// errors

obj = {
    name: "sai",
    age:20,
    // year: obj.age
}
console.log(obj) //reference error

obj = {
    name: "sai",
    age:20,
    year: this.age
}
console.log(obj) //undefined


//forEach

arr = [1,2 ,3,4,5]

arr.forEach((ele) => console.log(ele))


// objects in array

arr = [{name:"sai", age: 20}, {name:"anand", age: 30}]
console.log(arr[0]["name"])
console.log(arr[1].age)


// Math object

console.log(Math.round(7.6)) //8
console.log(Math.ceil(7.6)) //8
console.log(Math.floor(7.6)) //7
console.log(Math.trunc(7.6347)) //7 removes decimal part

rand = Math.floor(Math.random() * (max-min+1)) + min







