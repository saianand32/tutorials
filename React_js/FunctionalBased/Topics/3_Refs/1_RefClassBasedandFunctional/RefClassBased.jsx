import React, { Component } from 'react'

export default class Refs extends Component {

    constructor(props){
        super(props)
        //1. 1st  method using React.createRef()
        this.inputRef = React.createRef() //this is most used way to create ref

        //2. 2nd method (callback refs)
        this.cbRef = null;
        this.setCbRef = (e) => {
            this.cbRef = e 
        }
    }

    componentDidMount(){
        this.inputRef.current.focus()

        //cbref
        if(this.cbRef){
            this.cbRef.focus()
        }
    }
    
  render() {
    return (
      <div>
        <input ref = {this.inputRef}/>
        <input ref = {this.setCbRef}/>
      </div>
    )
  }
}
