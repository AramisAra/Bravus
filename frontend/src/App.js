import React from 'react';
import AppRouter from './router/AppRouter';
import './styles/index.css';

function App() {
  return (
    <div className="flex flex-col min-h-screen bg-gray-900 text-white">
      <AppRouter />
    </div>
  );
}

export default App;
