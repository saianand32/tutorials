import React, { useState } from 'react'

export default function D_useState (){

    const [count, setCount] = useState<number>(12)

  return (
    <div>
        <Child  setCount={setCount}/>
    </div>
  )
}


// But whats the type of the setCount ? lets pass it as props


type childCompProps = {
    setCount: React.Dispatch<React.SetStateAction<number>> //this is the type of setCount 
    // dont worry hover on the setCount where its defined and copy this
} 

const Child = ({ setCount }: childCompProps) => {
    setCount(9) //can pass number
    // setCount("lio") //err as cant pass string
    return(
        <>

        </>
    )
}


// 2. Default value props
//no need to specify the type for default props
const DefaultProps = ({ age = 20 }) => {
    return(
        <></>
    )
}