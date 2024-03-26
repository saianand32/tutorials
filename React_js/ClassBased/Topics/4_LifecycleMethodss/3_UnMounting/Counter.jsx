import React, { Component } from "react";

export default class Counter extends Component {

    constructor(props) {
        super(props)

        this.state = {
            counter: 0
        }

        console.log("Constructor")

        this.increment = () => this.setState({ counter: this.state.counter + 1 })
        this.decrement = () => this.setState({ counter: this.state.counter - 1 })
    }


    // Only one unmounting lifecycle method componentWillUnmount
    componentWillUnmount() {
        console.log("componentWillUnmount")
    }


    render() {
        console.log("Render")
        return (
            <>
                <h1>{this.state.counter}</h1>
                <button onClick={this.increment}>Increment</button>
                <button onClick={this.decrement}>Decrement</button>
            </>
        )
    }
}
