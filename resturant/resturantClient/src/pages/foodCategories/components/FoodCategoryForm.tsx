import FoodCategoryStore from "../store"
import { useEffect, useState } from "react"

function FoodCategoryForm() {
    const [values, setValues] = useState({
        title: "",
        description: "",
        photos: "",
    })
    const foodStore = FoodCategoryStore()
    useEffect(() => {
        setValues({
            ...(foodStore.targetItem || {
                title: "",
                description: "",
                photos: ""
            }),
        })
    }, [foodStore.targetItem])
    async function onSubmitHandler(e) {
        e.preventDefault()
        try {
            const payload = {
                title: values.title,
                description: values.description,
                photos: values.photos,
                status: 1
            }
            if (foodStore.targetItem?.id) {
                await foodStore.update(foodStore.targetItem.id, {
                    title: values.title,
                    description: values.description,
                    photos: undefined,
                    status: 1
                })
            } else {
                await foodStore.create(payload)
            }
            setValues({
                title: "",
                description: "",
                photos: ""
            })
        } catch (error) {
        }
    }

    function onChangeHandler(key: string, value: string) {
        setValues({
            ...values,
            [key]: value
        })
    }
    return (
        <form className="horizontal" onSubmit={onSubmitHandler}>
            <div className="fomControll">
                <label htmlFor="title">Title</label>
                <input type="text" id="title" value={values.title} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("title", e.target.value)
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

export default FoodCategoryForm