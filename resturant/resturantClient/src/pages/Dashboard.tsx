import { Route, Routes } from "react-router-dom"
import Home from "./Home"
import Users from "./users/Users"
import Nav from "@src/components/nav/Nav"
import FoodsPage from "./foods/FoodsPage"

function Dashboard() {
    return (
        <div className="dashboard">
            <header>
                <Nav />
            </header>
            <div className="rootElement">
                <Routes>
                    <Route path="/foods" element={<FoodsPage />} />
                    <Route path="/users" element={<Users />} />
                    <Route path="/" element={<Home />} />
                </Routes>
            </div>
        </div>
    )
}

export default Dashboard