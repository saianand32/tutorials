import React, { Component } from 'react'

export default class MyComp extends Component {

    constructor() {
        super()
        this.state = {
            id: 99,
            marks: 390,
            gender: 'male'
        }
        
    }

    incrementBy3(){ //wrong only increments by1
        this.setState({...this.state, id: this.state.id+1 })
        this.setState({...this.state, id: this.state.id+1 })
        this.setState({...this.state, id: this.state.id+1 })
    }

    incrementBy3Correct(){ //correct increments by 3
        this.setState((prevState) => {return {...prevState, id: prevState.id+1 }})
        this.setState((prev) => {return {...prev, id: prev.id+1 }})
        this.setState((prev) => {return {...prev, id: prev.id+1 }})
    }


    render() {
        return (
            <>
                <div>Details</div>
                <h1>Id = {this.state.id}</h1>
                <h1>marks = {this.state.marks}</h1>
                <h1>Gender = {this.state.gender}</h1>

                <button onClick={() => this.incrementBy3()}>Click Me</button>
                <button onClick={() => this.incrementBy3Correct()}>Click Me</button>
            </>

        )
    }

}
