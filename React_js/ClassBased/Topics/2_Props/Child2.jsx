import React, { Component } from 'react'

export default class Child2 extends Component {
  render() {
    const { gender, marks } = this.props
    return (
      <div>
        <h1>this is Child2 gender={gender} marks={marks}</h1>
      </div>

    )
  }
}
