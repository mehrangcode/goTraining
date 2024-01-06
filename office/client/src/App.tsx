import Dashboard from './pages/Dashboard'
import { BrowserRouter } from 'react-router-dom'

function App() {

  return (
    <div className="rootElement">
      <BrowserRouter>
      <Dashboard />
      </BrowserRouter>
    </div>
  )
}

export default App
