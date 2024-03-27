import React, { useEffect } from 'react'

const MemoChild = ({name}) => {

    useEffect(() => {
        console.log("childUpdated")
    })
   
  return (
    <div>MemoParent</div>
  )
}

export default React.memo(MemoChild)

// Note: 
// --- React.memo() can be used to shallow check prev props with the component 
// similar to making a component pure component in class based approach 

// without using the React.memo  parentUpdated and childUpdated are printed every 1 sec
// but en i add the React.memo() only parentUpdated priinted every 1 sec childUpdated printed once as then shallow comparison of props done by React.memo which is a higher order component