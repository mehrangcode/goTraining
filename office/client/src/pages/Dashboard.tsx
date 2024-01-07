import { Route, Routes } from "react-router-dom"
import Home from "./Home"
import Users from "./users/Users"
import IncomeLettersPage from "./incomeLetters/IncomeLettersPage"
import IncomeLetterForm from "@src/components/IncomeLetters/IncomeLetterForm"

function Dashboard() {
    return (
        <div className="dashboard">
            <header>
                <nav className="nav">OFFICE</nav>
            </header>
            <div className="rootElement">
                <Routes>
                    <Route path="/users" element={<Users />} />
                    <Route path="/income" element={<IncomeLettersPage />} />
                    <Route path="/income/modify" element={<IncomeLetterForm />} />
                    <Route path="/income/modify/:id" element={<IncomeLetterForm />} />
                    <Route path="/" element={<Home />} />
                </Routes>
            </div>
        </div>
    )
}

export default Dashboard