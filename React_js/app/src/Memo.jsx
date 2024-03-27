import React, { useState, useEffect } from 'react'

const MemoParent = ({name}) => {

    useEffect(() => {
        console.log("childUpdated")
    })
   
  return (
    <div>MemoParent</div>
  )
}

export default (MemoParent)