import React, { Component } from 'react'
import ForwardRef from './ForwardRef'

export default class ForwardParent extends Component {
    //main aim is that when i click button in parent component the childs input tag must be clicked using refs forwarding

    constructor(props){
        super(props)
        this.inpRef = React.createRef()
    }

    handleClick =()=> {
        this.inpRef.current.focus()
    }

  render() {
    return (
      <div>
        <ForwardRef reff = {this.inpRef}/>
        <button onClick={this.handleClick}>Focus input</button>
      </div>
    )
  }
}
