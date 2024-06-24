
//1.  setInterval -------------------------------

const fun = () => {
    console.log("hii")
}

setInterval(fun,1000)


// 2. clearInterval --------------------------------
// an intervalId is provided with that you can stop the interval

let count = 0;
const fun2 = () => {
count++;
console.log(count)
if(count == 4) clearInterval(intervalId)
}
 const intervalId = setInterval(fun2,1000);