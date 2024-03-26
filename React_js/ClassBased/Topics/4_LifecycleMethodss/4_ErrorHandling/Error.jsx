import React, { Component } from "react";

const obj = {};
const ErrorComp = () => <div>{obj.age.num}</div>;

class ErrorBoundary extends Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false };
  }
// ----------------------------------------------------------
  // Error handle Lifecycle methods

//   1. componentDidCatch
  componentDidCatch(error, errorInfo) {
    console.error("Error caught by ErrorBoundary:", error, errorInfo);
    this.setState({ hasError: true });
  }

// -------------------------------------------------------------

  render() {
    if (this.state.hasError) {
      return <h1>Something went wrong.</h1>;
    }
    return this.props.children;
  }
}

export default class Counter extends Component {
  constructor(props) {
    super(props);
    this.state = {
      counter: 0
    };

    this.increment = () => this.setState({ counter: this.state.counter + 1 });
    this.decrement = () => this.setState({ counter: this.state.counter - 1 });
  }

  render() {
    return (
      <ErrorBoundary>
        <>
          <h1>{this.state.counter}</h1>
          <button onClick={this.increment}>Increment</button>
          <button onClick={this.decrement}>Decrement</button>
          <ErrorComp />
        </>
      </ErrorBoundary>
    );
  }
}
