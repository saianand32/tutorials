const arr = [1,2,3,4,5]

// Map polyfils using Array.prototype
Array.prototype.myMap = function (callback) {
    const newArr = [];
    for (let e of this) {
      newArr.push(callback(e));
    }
    return newArr;
}

  
  function squareCallback (x) {
      return x*x;
  }
  
  const mappedArray = arr.myMap((e) => {
    return 2*e;
  });
  console.log(mappedArray)


  // make for filter and reduce on own