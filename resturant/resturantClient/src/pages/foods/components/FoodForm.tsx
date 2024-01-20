import Select from "@src/components/shared/select/Select"
import FoodCategoryStore from "@src/pages/foodCategories/store"
import FoodStore, { FoodType } from "@src/pages/foods/store"
import { useEffect, useState } from "react"

const defaultValue: FoodType = {
    name: "",
    description: "",
    photos: "",
    status: 1,
    categories: []
}
function FoodForm() {
    const [values, setValues] = useState<FoodType>(JSON.parse(JSON.stringify(defaultValue)))
    const foodStore = FoodStore()
    const categoriesStore = FoodCategoryStore()
    useEffect(() => {
        categoriesStore.fetchList()
    }, [])
    useEffect(() => {
        let convertedData: FoodType = JSON.parse(JSON.stringify(defaultValue))
        if (foodStore.targetItem) {
            convertedData = {
                ...foodStore.targetItem,
                categories: foodStore.targetItem.categories ? foodStore.targetItem.categories.map(x => x.id) : []
            }
        }
        setValues(convertedData)
    }, [foodStore.targetItem])
    console.log("V: ", values.categories)
    async function onSubmitHandler(e) {
        e.preventDefault()
        try {
            const payload: FoodType = {
                name: values.name,
                description: values.description,
                photos: values.photos,
                status: 1,
                categories: values.categories
            }
            if (foodStore.targetItem?.id) {
                await foodStore.update(foodStore.targetItem.id, payload)
            } else {
                await foodStore.create(payload)
            }
            setValues(JSON.parse(JSON.stringify(defaultValue)))
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
        <form className="horizontal" onSubmit={onSubmitHandler}>
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
            <div className="fomControll">
                <label htmlFor="description">description</label>
                <Select
                    mode="multiple"
                    value={values.categories as string[]}
                    onChange={value => setValues({
                        ...values,
                        categories: [...(values.categories as string[] || []), value]
                    })}
                    options={categoriesStore.list.map(record => ({ id: record.id, label: record.title }))}
                />

            </div>
            <button type="submit">{foodStore.targetItem ? "Edit" : "Save"}</button>
            {foodStore.targetItem ? <button type="button" onClick={() => foodStore.selectFood(undefined)}>Reset</button> : null}
        </form>
    )
}

export default FoodForm