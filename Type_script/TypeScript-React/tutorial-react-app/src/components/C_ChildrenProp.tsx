import React, { Children } from 'react'



// 1. React.ReactNode is like any ... can be a jsx element or string or bool any type
type compProps = {
    children: React.ReactNode
}

const C_ChildrenProp = ({ children }: compProps) => {
    return (
        <div>{children}</div>
    )
}


// in App call it like 

const App = () => {
    return (
        <>
            <C_ChildrenProp>Hey</C_ChildrenProp>  {/* can pass text */}
            <C_ChildrenProp><h1>Yo</h1></C_ChildrenProp> {/* can pass jsx element*/}
        </>
    )
}


// 2. if u want to pass only jsx Elements then do the compProps as below 
type compProps2 = {
    children: React.JSX.Element //this will allow children only to be a jsx  element
}



export default C_ChildrenProp