import React from 'react'


// mainly what is portal is that we can attach a component to any other div in index.html other than root
// with help of React portals

// step1. create another div in index.html and give id="portal1"
const Portal = () => {
  return ReactDom.createPortal (
    <h1>Portal</h1> ,
    document.getElementById("portal1")
  )
}

// step2. Now u can add this to app.js but this h1 tag will attach to the div with the id ="portal1" and not the id="root"

//NOte: Although portals are separate dom node not under root still all other things remain same .. u can pass props to children etc nothing else changes except the 
//fact that they are not under root node
export default Portal