import React from 'react'
import Hoc from './Hoc'

const B = ({name, age}) => {
  return (
   <>
     <div>{name}</div>
    <div>{age}</div>
   </>
  )
}

export default Hoc(B)