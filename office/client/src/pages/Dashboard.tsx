import { Route, Routes } from "react-router-dom"
import Home from "./Home"
import Users from "./users/Users"

function Dashboard() {
    return (
        <div className="dashboard">
            <header>
                <nav className="nav">OFFICE</nav>
            </header>
            <div className="rootElement">
                <Routes>
                    <Route path="/users" element={<Users />} />
                    <Route path="/" element={<Home />} />
                </Routes>
            </div>
        </div>
    )
}

export default Dashboard