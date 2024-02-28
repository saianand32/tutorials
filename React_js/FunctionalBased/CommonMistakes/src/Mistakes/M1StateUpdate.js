import React, { useState } from 'react'

const M1StateUpdate = () => {

    const [count, setCount] = useState(0)

    const handleClick = () => {

            // now u will  expect count to increment by 4 but that wonnt  
            //will only increase by 1
            // as setCount is not immediate
            // setCount(count+1)
            // setCount(count+1)
            // setCount(count+1)
            // setCount(count+1)

            // can do the below by using a callback with prev value

            setCount((prev) => prev+1)
            setCount((prev) => prev+1)
            setCount((prev) => prev+1)
            setCount((prev) => prev+1)
            //now the val increases by 4
    }

    return (
        <div>
            <h1>{count}</h1>
            <button style={{ width: "300px", height: "100px" }} onClick={handleClick}>Increment Count</button>
        </div>
    )
}

export default M1StateUpdate