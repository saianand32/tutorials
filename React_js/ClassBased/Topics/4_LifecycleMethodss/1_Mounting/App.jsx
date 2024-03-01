import React, { Component } from 'react'
import Counter from './Counter'

export default class App extends Component {

    constructor(props){
        super(props)
        this.state = {
            isMounted : false
        }

        this.mount = () => {this.setState({...this.state, isMounted: true })}
        this.unMount = () => {this.setState({...this.state, isMounted: false })}
    }


  render() {
    return (
      <div>
        <button onClick={this.mount}>Mount Counter component</button>
        <button onClick={this.unMount}>unMount Counnter component</button>

        {
            this.state.isMounted && <Counter/>
        }

      </div>
    )
  }
}

// Notes 
// 1. when i unmount component counter u can see componentWillUnmount is called
