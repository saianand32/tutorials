import React, { Component } from 'react';

class ErrorBoundary extends Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false };
  }

// -----------------------------------------------
  //Error handle Lifecycle methods

// 2.  getDerivedStateFromError
  static getDerivedStateFromError(error) {
    // Update state to indicate error
    return { hasError: true };
  }

//   ------------------------------------------
  render() {
    if (this.state.hasError) {
      // Render fallback UI when an error occurs
      return <h1>Something went wrong.</h1>;
    }
    // Render children as normal if no error occurred
    return this.props.children;
  }
}

class Example extends Component {
  constructor(props) {
    super(props);
    this.state = { counter: 0 };
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {
    // Simulate an error by dividing by zero
    this.setState(({ counter }) => ({ counter: 10 / counter }));
  }

  render() {
    return (
      <div>
        <h1>Counter: {this.state.counter}</h1>
        <button onClick={this.handleClick}>Divide by zero</button>
      </div>
    );
  }
}

function App() {
  return (
    <ErrorBoundary>
      <Example />
    </ErrorBoundary>
  );
}

export default App;
