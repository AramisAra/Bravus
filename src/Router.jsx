import React from "react";
import {Routes, Route, Navigate } from 'react-router-dom';
import Home from "./pages/Home";

function AppRoutes() {

    return (
    <>
        <main>
            <Routes>
                <Route path='/' element={<Home/>} />
            </Routes>
        </main>
    </>
    )
};

export default AppRoutes;