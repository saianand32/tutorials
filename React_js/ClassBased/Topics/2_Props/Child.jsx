import React, { Component } from "react";


class Child extends Component{
    render(){
        return (
            <>
                <h1>{name}</h1>
                <h1>{age}</h1>
                <button onClick={() => setNameAge("anand", 8)}>Click Me</button>

                <br />
                <br />
                <br />

                {this.props.children}
            </>
        )
    }
}


export default Child