import React from 'react'

// 1. 1st way 

// Note: below u can see React.JSX.Element type is optional to mention as typescript intellisense auto puts it anyways

const A_TypingProps = (props: {color: string, age?: number}): React.JSX.Element => {
  return (
    <>
        <div>{props.color}</div>
        <div>{props.age}</div>
    </>
  )
}


// 2. 2nd way to destructure 

const A_TypingProps2 = (props: {color: string, age: number}) => {
    const {color, age} = props
    return (
      <>
          <div>{color}</div>
          <div>{age}</div>
      </>
    )
  }


// 3. 3rd way to destructure first itself

const A_TypingProps3 = ({color, age}: {color: string, age: number}) => {
    return (
      <>
          <div>{color}</div>
          <div>{age}</div>
      </>
    )
  }



//   4. 4th way create a separate type to define the props (best way)

type compProps =  {
    color: string, 
    age?: number //optional field
}

const A_TypingProps4 = ({color, age}: compProps) => {
    return (
      <>
          <div>{color}</div>
          <div>{age}</div>
      </>
    )
  }


  // Note u can also use an interface but prefer types

  interface compProps5 {
    color: string, 
    age?: number //optional field
}

const A_TypingProps5 = ({color, age}: compProps5) => {
    return (
      <>
          <div>{color}</div>
          <div>{age}</div>
      </>
    )
  }


export default A_TypingProps