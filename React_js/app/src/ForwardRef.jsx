import React from 'react'


const ForwardRefChild = ((props) => {
    const {reff} = props
    return (
        <div>
            <input type="text" name="" id="" ref={reff} />
        </div>
      )
})


export default ForwardRefChild
