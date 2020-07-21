import React from 'react';
import logo from './logo.svg';
import './App.css';

type LightProps = {
  lightNumber: number;
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <span
        className="droppable">
        {LightButton({lightNumber: 1})}
        {LightButton({lightNumber: 2})}
        {LightButton({lightNumber: 3})}
        {LightButton({lightNumber: 4})}
        </span>
        {/* <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a> */}
      </header>
    </div>
  );
}


function LightButton(props: LightProps): any {

  function action(): any {
    alert(`Button for light ${props.lightNumber} pressed.`);
  }
  
  return (
    <div
    className="draggable">
    <button 
    onClick={action}
    
    >
      {`Light ${props.lightNumber}`}
    </button>
    </div>
  )
}

export default App;
