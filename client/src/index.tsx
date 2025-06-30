import '@assets/resets.css';

import React, { StrictMode } from 'react';
import ReactDOM from 'react-dom/client';

import App from './App';

const rootEl = document.getElementById('root') as HTMLElement;

if (rootEl) {
  const root = ReactDOM.createRoot(rootEl);

  root.render(
    <StrictMode>
      <App />
    </StrictMode>
  );
}
