import React from 'react';
import logo from './logo.svg';
import './App.css';
import A_TypingProps from './components/A_TypingProps';

function fun(age: number, name: string): number{
  return 0
}

function App() {
  return (
    <div className="App">
     <A_TypingProps color="anand" age={90}/>
    </div>
  );
}

export default App;
