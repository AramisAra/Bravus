import React from 'react';
import { Route, Routes, Navigate } from 'react-router-dom';
import Home from '../pages/Home'

function MainRouter() {
  return (
    <>
      <main>
        <Routes>
          <Route path="/" element={<Home/>} />
          <Route path="/home" element={<Home/>} />
        </Routes>
      </main>
    </>
  )
  
}

export default MainRouter;
