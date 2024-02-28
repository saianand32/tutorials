import React, { Component } from 'react'

export default class MyComp extends Component {

    constructor() {
        super()
        this.state = {
            id: 99,
            marks: 390,
            gender: 'male'
        }

        //need to do this why? see line  34
        this.handleClick1 = this.handleClick1.bind(this);
    }

    handleClick1(){
        this.setState({...this.state, id: 1 })
    }

    handleClick2(){
        this.setState({...this.state, id:2 })
    }

    handleClick3(){
        this.setState({...this.state, id:3 })
    }

    handleClick4 = () => {
        this.setState({...this.state, id:4 })
    }

    render() {
        const { id, marks, gender } = this.state
        return (
            <>
                <div>Details</div>
                <h1>Id = {id}</h1>
                <h1>marks = {marks}</h1>
                <h1>Gender = {gender}</h1>

                {/* 3 ways  */}

                {/* 1.in this way u need to add bind method in the constructor as in line 12 to access this.setState in the function */}
                <button onClick={this.handleClick1}>Click Me</button> 

                {/* 2. Can use bind in the atribute itself  */}
                <button onClick={this.handleClick2.bind(this)}>Click Me</button> 

                {/* 3. use an arrow function and return the method (no need to bind) */}
                <button onClick={() => this.handleClick3()}>Click Me</button> {/* No need to add a bind method */}

                {/* 4. Easiest way -- use arrow function to define method then simply use it no need to bind */}
                <button onClick={this.handleClick4}>Click Me</button> 

            </>

        )
    }

}
