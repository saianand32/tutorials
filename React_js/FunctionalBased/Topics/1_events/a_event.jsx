import React, { useRef } from 'react';

const MyComponent = () => {
  const buttonRef = useRef(null);

  const handleClick = (e) => {
    e.stopPropagation(); // used to prevent onClick bubble up the dom tree (event bubbling)
    e.preventDefault(); //used to prevent default behaviour especially in forms
    console.log('Button clicked');
  };

  return (
    <div onClick={() => console.log('Div clicked')}>
      <button ref={buttonRef} onClick={handleClick}>Click me</button>
    </div>
  );
};

export default MyComponent;
