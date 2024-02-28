import React, { useState, useEffect } from 'react';

const EventListenerDemo = () => {
  const [message, setMessage] = useState('');
  const [inputValue, setInputValue] = useState('');

  useEffect(() => {
    console.log(message)
  }, [message])

  const handleMouseEnter = () => {
    setMessage('Mouse entered');
  };

  const handleMouseLeave = () => {
    setMessage('Mouse left');
  };

  const handleKeyDown = (e) => {
    setMessage(`Key pressed: ${e.key}`);

    //using the enter key to submit data
    if(e.key == 'Enter') alert("submit the data")
  };

  const handleKeyUp = () => {
    setMessage('Key released');
  };

  const handleClick = () => {
    setMessage('Mouse clicked');
  };

  const handleFocus = () => {
    setMessage('Input field focused');
  };

  const handleChange = (e) => {
    setInputValue(e.target.value);
    setMessage(`Input value changed: ${e.target.value}`);
  };

  return (
    <div>
      <p>Move your mouse over here:</p>
      <div
        style={{ width: '200px', height: '100px', backgroundColor: 'lightgray' }}
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}
        onMouseDown={handleClick}
      >
        {message}
      </div>
      <p>Press a key:</p>
      <input
        type="text"
        onKeyDown={handleKeyDown}
        onKeyUp={handleKeyUp}
        onFocus={handleFocus}
        onChange={handleChange}
        value={inputValue}
        onCopy={() => console.log("copied")}
        onCut={() => console.log("cut")}
        onPaste={() => console.log("pasted")}
      />
    </div>
  );
};

export default EventListenerDemo;
