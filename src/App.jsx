import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import AppRoutes from './Router'

export default function App() {
  return (
    <div className="flex min-h-screen bg-main-black">
      <AppRoutes/>
    </div>
  )
}
