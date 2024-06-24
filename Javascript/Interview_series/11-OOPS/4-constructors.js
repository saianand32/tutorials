
// using a constructor function I can define multiple objects with same properties

function car (brand, model, color){
    this.brand = brand;
    this.model = model;
    this.color = color;
}

let car1 = new car('BMW',"X-5","white")
console.log(car1)

// U can also put functions as obj properties via constructor function

function car2 (brand, model, color){
    this.brand = brand;
    this.model = model;
    this.color = color;

    this.drive = function(){
        console.log("Im driving "+ this.brand)
    }
}

let myCar = new car2('Ferari',"wednoih","white")
myCar.drive()