  

// 1.  Generic classes

class List<T>{
    arr: T[]

    constructor(){
        this.arr = []
    }

    add(val: T){
        this.arr.push(val)
    }

}

let al = new List<string>();

al.add("sai")
al.add("anand")
al.add(9) //err as u cant push a number