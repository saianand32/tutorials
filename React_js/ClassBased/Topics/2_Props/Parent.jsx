import React, { Component } from "react";
import Child2 from "./Child2";
import Child from "./Child";


class Parent extends Component{

    constructor(){
        super()
        this.state = {
            name: "sai",
            age: 22
        }
    }

    setNameAge = (name, age) => { //method to be passed as a prop
            this.setState((prevState) =>  {return {...prevState, name, age}})
    }

    render(){
        return (
            <>  
                <Child setNameAge={this.setNameAge} name={this.state.name} age = {this.state.age} >
                    <Child2 gender="male" marks={88}/> {/* See child2 for destructuring props */}
                </Child>
            </>
        )
    }
}


export default Parent