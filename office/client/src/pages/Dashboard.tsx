import { Route, Routes } from "react-router-dom"
import Home from "./Home"
import Users from "./users/Users"

function Dashboard() {
    return (
        <div>
            <div className="nav">Nav</div>
            <Routes>
                <Route path="/users" element={<Users />} />
                <Route path="/" element={<Home />} />

            </Routes>
        </div>
    )
}

export default Dashboard