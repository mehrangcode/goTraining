import { useEffect } from "react"
import FoodStore from "./store"
import FoodForm from "@src/components/foods/FoodForm"

function FoodsPage() {
  const foodStore = FoodStore()
  useEffect(() => {
    foodStore.fetchList()
  }, [])
  return (
    <div className="foodPage">
      <FoodForm />
      <table className="list">
        <thead>
          <tr>
            <th>Name</th>
            <th>Description</th>
            <th>Status</th>
            <th>action</th>
          </tr>
        </thead>
        <tbody>

          {foodStore.list?.map(record => {
            return <tr>
              <td>{record.name}</td>
              <td>{record.description}</td>
              <td>{record.status}</td>
              <td>
                <div className="action">
                  <button onClick={() => foodStore.changeStatus(record.id, record.status === 1 ? 0 : 1)}>ChangeStatus</button>
                  <button onClick={() => foodStore.selectFood(record)}>Edit</button>
                  <button onClick={() => foodStore.delete(record.id)}>Delete</button>
                </div>
              </td>
            </tr>
          })}
        </tbody>
      </table>
    </div>
  )
}

export default FoodsPage