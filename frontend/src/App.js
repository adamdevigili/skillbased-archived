import React from 'react';
import './App.css';
import SportsDropdown from './components/Dropdown';
import GoButton from './components/GoButton';
import HomepageLayout from "./components/HomepageLayout";

function App() {
  return (
      <div className="centered">
        <div className="title">
          <h1>SKILLBASED.IO</h1>
        </div>
        <div className="sport_dropdown">
          <h3>Select a sport to get started</h3>
          <SportsDropdown />
          <div className="go_button">
            <GoButton />
          </div>
        </div>
      </div>
  );
}

export default App;
