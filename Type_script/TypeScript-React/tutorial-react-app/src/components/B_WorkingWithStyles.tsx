import React from 'react'

//React gives special type for styles React.CSSProperties

type CompProps = {
    style: React.CSSProperties, //for css
}

const B_WorkingWithStyles = ({style}: CompProps) => {
  return (
    <div style={style}> B_WorkingWithStyles</div>
  )
}



// in App.tsx call the <B_WorkingWithStyles/>

export default B_WorkingWithStyles

const App = () => {
    return (
      <div>
        <B_WorkingWithStyles 
            style={{
                color:"white",
                fontSize:10,
                borderRadius: 8
            }}
        />
      </div>
    )
  }