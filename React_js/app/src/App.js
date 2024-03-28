import { useState, useEffect, Component } from 'react';
import './App.css';
import Memo from './Memo';
import Refs from './Refs';
import RefsFunctional from './RefsFunctional';
import ForwardParent from './ForwardParent';
import ForwardRef from './ForwardRef';
import ErrorBoundary from './ErrorBoundary';


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
        <ErrorBoundary>
          <ForwardRef heroName={"saishwar"} />
          <ForwardRef heroName={"Anand"} />
          <ForwardRef heroName={"Joker"} />
        </ErrorBoundary>
      </>
    )
  }
}

export default App;
