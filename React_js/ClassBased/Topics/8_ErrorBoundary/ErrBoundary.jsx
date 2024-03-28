import React, { Component } from 'react'

export default class ErrorBoundary extends Component {

    constructor(props) {
      super(props)
    
      this.state = {
         hasErr: false
      }
    }

    static getDerivedStateFromError(error){
        return {
            hasErr: true
        }
    }

  render() {
    if(this.state.hasErr) return <h1>Some Error</h1>
    
    else return this.props.children
  }
}
