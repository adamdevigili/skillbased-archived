import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import APIApp from './APIApp';
import * as serviceWorker from './serviceWorker';

let host = window.location.host
const parsedData = host.split(".");

if (parsedData[0] === "api") {
    const subDomain = parsedData[0];
    ReactDOM.render(
        <React.StrictMode>
            <APIApp />
        </React.StrictMode>,
        document.getElementById('root')
    );
} else {
    ReactDOM.render(
        <React.StrictMode>
            <App />
        </React.StrictMode>,
        document.getElementById('root')
    );
}

// ReactDOM.render(
//   <React.StrictMode>
//     <App />
//   </React.StrictMode>,
//   document.getElementById('root')
// );

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
