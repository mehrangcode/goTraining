import FoodStore from "@src/pages/foods/store"
import { useEffect, useState } from "react"

function FoodForm() {
    const [values, setValues] = useState({
        name: "",
        description: "",
        photos: "",
    })
    const foodStore = FoodStore()
    useEffect(() => {
        setValues({
            ...(foodStore.targetItem || {
                name: "",
                description: "",
                photos: ""
            }),
        })
    }, [foodStore.targetItem])
    async function onSubmitHandler(e) {
        e.preventDefault()
        try {
            const payload = {
                name: values.name,
                description: values.description,
                photos: values.photos,
                status: 1
            }
            if (foodStore.targetItem?.id) {
                await foodStore.update(foodStore.targetItem.id, {
                    name: values.name,
                    description: values.description,
                    photos: undefined,
                    status: 1
                })
            } else {
                await foodStore.create(payload)
            }
            setValues({
                name: "",
                description: "",
                photos: ""
            })
        } catch (error) {
        }
    }

    function onChangeHandler(name: string, value: string) {
        setValues({
            ...values,
            [name]: value
        })
    }
    return (
        <form className="foodForm" onSubmit={onSubmitHandler}>
            <div className="fomControll">
                <label htmlFor="name">Name</label>
                <input type="text" id="name" value={values.name} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("name", e.target.value)
                }} />
            </div>
            <div className="fomControll">
                <label htmlFor="description">description</label>
                <input type="text" id="description" value={values.description} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("description", e.target.value)
                }} />
            </div>
            <button type="submit">{foodStore.targetItem ? "Edit" : "Save"}</button>
            {foodStore.targetItem ? <button type="button" onClick={() => foodStore.selectFood(undefined)}>Reset</button> : null}
        </form>
    )
}

export default FoodForm