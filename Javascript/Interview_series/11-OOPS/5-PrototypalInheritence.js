
let person1 = {
    name:"Adam",
    age:25,

    showDetails: function(){
        console.log(this.name+" "+this.age)
    }
}


let person2 = {
    name:"Steve",
}

person2.__proto__ = person1;

person2.showDetails() // even though person2 doesnt have shoDetails
                      // it accesses it from person1 this is prototypal inheritance