import React from 'react'

const MyComponent = ({heroName}) => {
    if(heroName === 'Joker') throw new Error('Not a hero')

  return (
    <h1>{heroName}</h1>
  )
}

export default MyComponent