import React, { Component, useEffect } from "react";

export default class Counter extends Component {

    // 1st mounting lifecycle method
    // place where u define state and call super(props) to attach props to this 
    constructor(props) {
        super(props)

        this.state = {
            counter: 0
        }

        this.increment = () => this.setState({ counter: this.state.counter + 1 })
        this.decrement = () => this.setState({ counter: this.state.counter - 1 })
    }


    //UPDATE METHODS  (getDerivedStateFromProps, render, getSnapShotBeforeUpdate, shouldComponentUpdate, componentDidUpdate)

    static getDerivedStateFromProps(props, state) {
        console.log("getDerivedStateFromProps")
        return { ...state, age: props.age }
        // return null  if no change in state
    }

    shouldComponentUpdate(nextProps, nextState){ // no ajax calls here as its called frequently
        //check if nextProps or nextState is changed
        return true; //or false on check basis
    }

    getSnapshotBeforeUpdate(prevProps, prevState){ //rarely used
        //called right before changes from virtual dom to be reflected to real dom
        // for ex can capture users scroll positions and maintain that after update 
        return "hello sai"
    }

    componentDidUpdate(prevProps, prevState, snapshot) { // use for ajax calls etc 
        console.log("componentDidUpdate")
        console.log(snapshot) //prints hello sai as recieved a snapshot from getSnapshotBeforeUpdate
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
//  getDerivedStateFromProps, shouldComponentUpdate, getSnapshotBeforeUpdate, render and componentDidUpdate
