import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <section className='bg-emerald-500'>
      <App />
    </section>
  </React.StrictMode>,
)
