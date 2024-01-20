import { DeleteOutlined, EditOutlined } from "@ant-design/icons"
import FoodCategoryStore, { FoodCategoryType } from "../store"

interface IProps {
    food: FoodCategoryType
}
function FoodCard({
    food
}: IProps) {
    const foodCategoryStore = FoodCategoryStore()
    return (
        <div className="itemCard">
            <div className="itemAvatar">A</div>
            <div className="itemName">{food.title}</div>
            <small className="itemDescription">{food.description}</small>
            <div className="itemActionWrapper">
                <EditOutlined onClick={() => foodCategoryStore.selectFood(food)}/>
                <DeleteOutlined onClick={() => foodCategoryStore.delete(food.id)}/>
                {/* <button onClick={() => foodStore.changeStatus(food.id, food.status === 1 ? 0 : 1)}>ChangeStatus</button> */}
            </div>
        </div>
    )
}

export default FoodCard