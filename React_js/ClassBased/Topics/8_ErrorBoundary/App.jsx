import { Component } from 'react';
import './App.css';
import MyComponent from './MyComponent';
import ErrorBoundary from './ErrorBoundary';


class App extends Component {

  constructor(props) {
    super(props)
    this.state = {
      name: "sai"
    }
  }

  // if anywhere error thrown app wont crash and the error handling method we define will be implemented
  render() {
    return (
      <>
        <ErrorBoundary>
          <MyComponent heroName={"saishwar"} />
          <MyComponent heroName={"Anand"} />
          <MyComponent heroName={"Joker"} />
        </ErrorBoundary>
      </>
    )
  }
}

export default App;
