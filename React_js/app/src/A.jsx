import React from 'react'
import Hoc from './Hoc'

const A = ({name, age}) => {
  return (
    <>
        <div>{name}</div>
    <div>{age}</div>
    </>
  )
}

export default Hoc(A)