import { Route, Routes } from "react-router-dom"
import Home from "./Home"
import Users from "./users/Users"
import Nav from "@src/components/nav/Nav"

function Dashboard() {
    return (
        <div className="dashboard">
            <header>
                <Nav />
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