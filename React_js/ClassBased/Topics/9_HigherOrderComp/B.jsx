import React from 'react'
import Hoc from './Hoc'

const B = ({name}) => {
  return (
    <div>{name}</div>
  )
}

export default Hoc(B)