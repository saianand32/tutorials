import React from 'react'
import './Border.css'

const Border = ({ children }) => {
  return (
    <div style={{border: " 2px solid orange", width: "100%", height:"100vh"}}>
        {children}
    </div>
  )
}

export default Border