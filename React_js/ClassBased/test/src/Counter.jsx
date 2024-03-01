import React, { Component, useEffect } from "react";

export default class Counter extends Component {

    // 1st mounting lifecycle method
    // place where u define state and call super(props) to attach props to this 
    constructor(props) {
        super(props)

        this.state = {
            counter: 0
        }

        console.log("Constructor")

        this.increment = () => this.setState({ counter: this.state.counter + 1 })
        this.decrement = () => this.setState({ counter: this.state.counter - 1 })
    }

    static getDerivedStateFromProps(props, state) {
        console.log("getDerivedStateFromProps")
        return { ...state, age: props.age }
    }

    componentDidMount() {
        console.log("componentDidMount")
    }


    componentWillUnmount() {
        console.log("componentWillUnmount")
    }


    // next topic method but just for demo
    componentDidUpdate() {
        console.log("componentDidUpdate")
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

// Notes 
// 1. order of execution 
// constructor, getDerivedStateFromProps, render, componentDidMount 

// 2. when i click increment decrement functions render method is called and componentDidUpdate is called. componentDidMount not called is called only once 
// so componentDidMount is like useEffect(() => {}, []) 
// so componentDidUpdate is like useEffect(() => {}) 
// so componentShouldUpdate is like useEffect(() => {}, [state1, state2..etc]) 

// 3. when component unmounts componentWillUnmount is called 