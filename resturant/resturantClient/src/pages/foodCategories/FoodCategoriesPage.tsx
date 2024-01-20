import { useEffect } from "react"
import FoodCategoryStore from "./store"
import FoodCategoryForm from "./components/FoodCategoryForm"
import FoodCard from "./components/FoodCard"
import './style.css'
function FoodCategoriesPage() {
  const foodCategoryStore = FoodCategoryStore()
  useEffect(() => {
    foodCategoryStore.fetchList()
  }, [])
  return (
    <div className="foodCategoriesPage">
      <FoodCategoryForm />
      <div className="itemsListPreview">
        {foodCategoryStore.list?.map(food => {
          return <FoodCard food={food} />
        })}
      </div>
    </div>
  )
}

export default FoodCategoriesPage