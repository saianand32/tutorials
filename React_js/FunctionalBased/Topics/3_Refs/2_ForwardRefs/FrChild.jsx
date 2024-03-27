import React from 'react'

//use React.ForwardRef // also u can just pass ref as a prop as well but this is standard
const ForwardRefChild = React.forwardRef((props, ref) => {
    return (
        <div>
            <input type="text" name="" id="" ref={ref} />
        </div>
      )
})


export default ForwardRefChild
