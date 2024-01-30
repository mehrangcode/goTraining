import MenuStore, { MenuType } from "@src/pages/menus/store"
import { useEffect, useState } from "react"

const defaultValue: MenuType = {
    title: "",
    description: "",
    status: 1,
    sections: [],
}
function MenuForm() {
    const [values, setValues] = useState<MenuType>(JSON.parse(JSON.stringify(defaultValue)))
    const menuStore = MenuStore()
    useEffect(() => {
        setValues({
            ...(menuStore.targetItem || JSON.parse(JSON.stringify(defaultValue))),
        })
    }, [menuStore.targetItem])
    async function onSubmitHandler(e) {
        e.preventDefault()
        try {
            const payload: MenuType = { ...values,
            sections: [
                {
                    title: "starter",
                    description: "section description",
                    foods: [
                        {
                            food_id: "1",
                            price: 200000
                        },
                        {
                            food_id: "3",
                            price: 200000
                        }
                    ]
                },
                {
                    title: "main",
                    description: "section description",
                    foods: [
                        {
                            food_id: "1",
                            price: 80000
                        },
                        {
                            food_id: "2",
                            price: 500000
                        },
                        {
                            food_id: "1",
                            price: 1000000
                        }
                    ]
                }
            ]
            }
            if (menuStore.targetItem?.id) {
                await menuStore.update(menuStore.targetItem.id, payload)
            } else {
                await menuStore.create(payload)
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
        <form onSubmit={onSubmitHandler}>
        <button type="submit">{menuStore.targetItem ? "Edit" : "Save"}</button>
        {menuStore.targetItem ? <button type="button" onClick={() => menuStore.selectMenu(undefined)}>Close</button> : null}
            <div className="fomControll">
                <label htmlFor="title">title</label>
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
            {!menuStore.targetItem ? <div className="fomControll">
                <label htmlFor="status">status</label>
                <input type="text" id="status" value={values.status} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("status", e.target.value)
                }} />
            </div> : null}
            <div className="sectionsWrapper">
                
            </div>
        </form>
    )
}

export default MenuForm