import React, { PureComponent } from 'react'

export default class PureComp extends PureComponent { //here note that we extend PureComponent class of react
  render() {
    return (
      <div>PureComp</div>
    )
  }
}

//Mainly theory

// --- Normal components dont have a shallow check of prevstate and prev props in the shouldComponentUpdate method 
// --- On the other hand Pure components only return true in shouldComponentUpdate method if changes are there by doing shallow comparison 
