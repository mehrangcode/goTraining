import axios from 'axios';
import Dashboard from './pages/Dashboard'
import { BrowserRouter } from 'react-router-dom'

function App() {

  axios.interceptors.response.use(function (response) {
    return response;
  }, function (error) {
    console.error("axios error", error)
    return Promise.reject(error);
  });

  return (
    <div className="app">
      <BrowserRouter>
      <Dashboard />
      </BrowserRouter>
    </div>
  )
}

export default App
