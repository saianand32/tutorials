import { Component } from 'react';
import './App.css';
import MemoChild from './MemoChild';


class MemoParent extends Component{

  constructor(props){
    super(props)
      this.state = {
        name:"sai"
      }
  }

  componentDidMount(){
    setInterval(() => { //basically resetting state to same thing
      this.setState({
        name:"sai"
      })
  }, 1000)
  }

  componentDidUpdate(){
    console.log("Parentupdated")
  }


  render(){
    return (
      <>
         <MemoChild name={this.state.name}/>
      </>
    )
  }
}

export default MemoParent;
