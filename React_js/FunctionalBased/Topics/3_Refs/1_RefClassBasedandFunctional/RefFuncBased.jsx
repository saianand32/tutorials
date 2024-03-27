import React, { useEffect, useRef } from 'react'

const RefsFunctional = () => {

    const inpRef = useRef(null)

    useEffect(() => {
        inpRef.current.focus()
    }, [])

  return (
    <div>
        <input type="text" name="" id="" ref = {inpRef}/>
    </div>
  )
}

export default RefsFunctional