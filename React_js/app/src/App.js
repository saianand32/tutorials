import { useState, useEffect, Component } from 'react';
import './App.css';
import A from './A';
import B from './B';


class App extends Component {

  constructor(props) {
    super(props)
    this.state = {
      name: "sai"
    }
  }

  render() {
    return (
      <>
        <A age={90}/>
        <B/>
      </>
    )
  }
}

export default App;
