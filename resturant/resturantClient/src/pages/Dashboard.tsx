import { Route, Routes } from "react-router-dom"
import Home from "./Home"
import Users from "./users/Users"
import Nav from "@src/components/nav/Nav"
import FoodsPage from "./foods/FoodsPage"
import MenusPage from "./menus/MenusPage"
import FoodCategoriesPage from "./foodCategories/FoodCategoriesPage"
import TablesPage from "./tables/TablePage"

function Dashboard() {
    return (
        <div className="dashboard">
            <header>
                <Nav />
            </header>
            <div className="rootElement">
                <Routes>
                    <Route path="/tables" element={<TablesPage />} />
                    <Route path="/menus" element={<MenusPage />} />
                    <Route path="/foods/categories" element={<FoodCategoriesPage />} />
                    <Route path="/foods" element={<FoodsPage />} />
                    <Route path="/users" element={<Users />} />
                    <Route path="/" element={<Home />} />
                </Routes>
            </div>
        </div>
    )
}

export default Dashboard