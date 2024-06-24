const arr = [1, 2, 3,4,5,6];


// 1. Map -------------------------------------

function polymap (arr, callback){
  const newArr = [];
  for (let e of arr) {
    newArr.push(callback(e));
  }
  return newArr;
};

function squareCallback (x) {
    return x*x;
}

const mappedArray = polymap(arr,squareCallback);
console.log(mappedArray)


// 1. Filter -------------------------------------

function polyfilter (arr, callback){
    const newArr = []
    for(let e of arr){
        if(callback(e)){
            newArr.push(e);
        }
    }
    return newArr
}

function filterCallback (e){
if(e%2==0) return true;
return false;
}

const filteredarr = polyfilter(arr, filterCallback);
console.log(filteredarr)



// 3. Reduce -----------------------------------

function polyReduce(arr,accumulator, callback){
    let ans = accumulator;
    for(let e of arr){
        ans = callback(ans,e);
    }
    return ans;
}

function reduceCallback(accumulator, element){
    return accumulator + element;
}

const reducedArray = polyReduce(arr,0,reduceCallback);
console.log(reducedArray)


