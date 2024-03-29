import React from 'react'
import Hoc from './Hoc'

const A = ({name}) => {
  return (
    <div>{name}</div>
  )
}

export default Hoc(A)