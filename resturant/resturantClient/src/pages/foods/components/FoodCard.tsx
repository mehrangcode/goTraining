import { DeleteOutlined, EditOutlined } from "@ant-design/icons"
import FoodStore, { FoodType } from "../store"

interface IProps {
    food: FoodType
}
function FoodCard({
    food
}: IProps) {
    const foodStore = FoodStore()
    return (
        <div className="foodCard">
            <div className="foodAvatar">A</div>
            <div className="foodName">{food.name}</div>
            <small className="foodDescription">{food.description}</small>
            <div className="foodActionWrapper">
                <EditOutlined onClick={() => foodStore.selectFood(food)}/>
                <DeleteOutlined onClick={() => foodStore.delete(food.id)}/>
                {/* <button onClick={() => foodStore.changeStatus(food.id, food.status === 1 ? 0 : 1)}>ChangeStatus</button> */}
            </div>
        </div>
    )
}

export default FoodCard