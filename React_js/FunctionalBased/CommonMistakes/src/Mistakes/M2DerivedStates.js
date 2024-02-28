import React, { useEffect, useState } from 'react'


//************************Bad Way**********************************

// const M2DerivedStates = () => {

//     const COST = 100;

//     const [ quantity, setQuantity ] = useState(0)
//     const [ total, setTotal ] = useState(quantity*COST)


//     useEffect(() => {
//         setTotal(quantity*COST)
//     }, [quantity])


//     const handleClick = () => {
//         setQuantity((prev) => prev +1)
//     }


//     return (
//         <div>
//             <h1>Quantity: {quantity}</h1>
//             <h1>Total: {total}</h1>
//             <button style={{ width: "300px", height: "100px" }} onClick={handleClick}>Increment Count</button>
//         </div>
//     )
// }

//******************************* Good Way **********************************

const M2DerivedStates = () => {

    const COST = 73;

    const [ quantity, setQuantity ] = useState(0)
    const total = quantity*COST; 
    
    // since total is a derived state u can assign like this 
    // and its auto updated as wen u change quantity state component is rerendered
    // and so total gets updated .... the useEffect and another useState is unnecessary



    const handleClick = () => {
        setQuantity((prev) => prev +1)
    }


    return (
        <div>
            <h1>Quantity: {quantity}</h1>
            <h1>Total: {total}</h1>
            <button style={{ width: "300px", height: "100px" }} onClick={handleClick}>Increment Count</button>
        </div>
    )
}

export default M2DerivedStates