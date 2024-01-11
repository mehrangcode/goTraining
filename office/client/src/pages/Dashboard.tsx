import IncomeLetterForm from "@src/components/IncomeLetters/IncomeLetterForm"
import IssuedLetterForm from "@src/components/IssuedLetters/IssuedLetterForm"
import Nav from "@src/components/nav/Nav"
import { Route, Routes } from "react-router-dom"
import Home from "./Home"
import IncomeLettersPage from "./incomeLetters/IncomeLettersPage"
import IssuedLettersPage from "./issuedLetters/IssuedLettersPage"
import SubjectsPage from "./subjects/SubjectsPage"
import Users from "./users/Users"

function Dashboard() {
    return (
        <div className="dashboard">
            <header>
                <Nav />
            </header>
            <div className="rootElement">
                <Routes>
                    <Route path="/users" element={<Users />} />
                    <Route path="/letters/income" element={<IncomeLettersPage />} />
                    <Route path="/letters/income/modify" element={<IncomeLetterForm />} />
                    <Route path="/letters/income/modify/:letterId" element={<IncomeLetterForm />} />
                    <Route path="/letters/issued" element={<IssuedLettersPage />} />
                    <Route path="/letters/issued/modify" element={<IssuedLetterForm />} />
                    <Route path="/letters/issued/modify/:letterId" element={<IssuedLetterForm />} />
                    <Route path="/subjects" element={<SubjectsPage />} />
                    <Route path="/" element={<Home />} />
                </Routes>
            </div>
        </div>
    )
}

export default Dashboard