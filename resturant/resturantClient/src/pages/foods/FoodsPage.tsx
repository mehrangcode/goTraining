import { useEffect } from "react"
import FoodStore from "./store"
import FoodForm from "./components/FoodForm"
import FoodCard from "./components/FoodCard"
import './style.css'
function FoodsPage() {
  const foodStore = FoodStore()
  useEffect(() => {
    foodStore.fetchList()
  }, [])
  return (
    <div className="foodPage">
      <FoodForm />
      <div className="foodsListPreview">
        {foodStore.list?.map(food => {
          return <FoodCard food={food} />
        })}
      </div>
    </div>
  )
}

export default FoodsPage