import React, { Component } from "react";

export default class App extends Component {

    constructor() {
        super()
        this.state = {
            inputData: {}
        }
    }

    handleChange = (e) => {
        this.setState({ ...this.state, inputData: { ...this.state.inputData, [e.target.name]: e.target.value } })
    }

    handleSubmit = (e) => {
        e.preventDefault()
        console.log(this.state.inputData)
    }

    render() {
        return (
            <>
                <form action="" onSubmit={this.handleSubmit}>
                    <input type="text" name="name" id="" placeholder="name" onChange={this.handleChange} />
                    <input type="text" name="age" id="" placeholder="age" onChange={this.handleChange} />

                    <button type="submit">Submit</button>
                </form>
            </>
        )
    }
}