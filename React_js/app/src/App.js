import { useState, useEffect, Component } from 'react';
import './App.css';
import Memo from './Memo';
import Refs from './Refs';
import RefsFunctional from './RefsFunctional';
import ForwardParent from './ForwardParent';


class App extends Component{

  constructor(props){
    super(props)
      this.state = {
        name:"sai"
      }
  }

  componentDidMount(){
  //   setInterval(() => { //basically resetting state to same thing
  //     this.setState({
  //       name:"sai"
  //     })
  // }, 1000)
  }

  componentDidUpdate(){
    console.log("updated")
  }


  render(){
    return (
      <>
         {/* <RefsFunctional/> */}
         {/* <Refs/> */}
         <ForwardParent/>
      </>
    )
  }
}

export default App;
