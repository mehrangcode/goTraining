import { Route, Routes } from "react-router-dom"
import Home from "./Home"
import Users from "./users/Users"
import IncomeLettersPage from "./incomeLetters/IncomeLettersPage"
import IncomeLetterForm from "@src/components/IncomeLetters/IncomeLetterForm"
import SubjectsPage from "./subjects/SubjectsPage"
import { Link } from "react-router-dom"
import IssuedLettersPage from "./issuedLetters/IssuedLettersPage"
import IssuedLetterForm from "@src/components/IssuedLetters/IssuedLetterForm"

function Dashboard() {
    return (
        <div className="dashboard">
            <header>
                <nav className="nav"><Link to={"./"}>OFFICE</Link></nav>
            </header>
            <div className="rootElement">
                <Routes>
                    <Route path="/users" element={<Users />} />
                    <Route path="/letters/income" element={<IncomeLettersPage />} />
                    <Route path="/letters/income/modify" element={<IncomeLetterForm />} />
                    <Route path="/letters/income/modify/:id" element={<IncomeLetterForm />} />
                    <Route path="/letters/issued" element={<IssuedLettersPage />} />
                    <Route path="/letters/issued/modify" element={<IssuedLetterForm />} />
                    <Route path="/letters/issued/modify/:id" element={<IssuedLetterForm />} />
                    <Route path="/subjects" element={<SubjectsPage />} />
                    <Route path="/" element={<Home />} />
                </Routes>
            </div>
        </div>
    )
}

export default Dashboard