import { useState, useEffect } from "react"
import MenuStore, { MenuType } from "./store"
import Select from "@src/components/shared/select/Select"
import FoodStore from "../foods/store"


const defaultValue: MenuType = {
  title: "",
  description: "",
  status: 1,
  sections: [],
}
function MenuFormPage() {
  const [values, setValues] = useState<MenuType>(JSON.parse(JSON.stringify(defaultValue)))
  const menuStore = MenuStore()
  const foodStore = FoodStore()
  useEffect(() => {
    foodStore.fetchList()
  }, [])
  useEffect(() => {
    setValues({
      ...(menuStore.targetItem || JSON.parse(JSON.stringify(defaultValue))),
    })
  }, [menuStore.targetItem])
  async function onSubmitHandler(e) {
    e.preventDefault()
    try {
      const payload: MenuType = {
        ...values,
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
      <p>Sections <button onClick={() => {
        setValues({
          ...values,
          sections: [...values.sections,
          {
            title: "",
            description: "",
            foods: [
              {
                food_id: null,
                price: null
              }
            ]
          }
          ]
        })
      }}>Add New Section</button> </p>
      <div className="sectionsWrapper">
        {values.sections.map((section, sectionIndex) => {
          return <section className="menuSection">
            <div className="fomControll">
              <label htmlFor="title">title</label>
              <input type="text" id="title" value={section.title} onChange={(e) => {
                e.preventDefault()
                const updatedData = JSON.parse(JSON.stringify(values))
                updatedData.sections = updatedData.sections.map((item, itemIndex) => {
                  if (itemIndex === sectionIndex) {
                    return { ...item, title: e.target.value }
                  }
                  return item
                })
                setValues(updatedData)
              }} />
            </div>
            <div className="fomControll">
              <label htmlFor="description">description</label>
              <input type="text" id="description" value={section.description} onChange={(e) => {
                e.preventDefault()
                const updatedData = JSON.parse(JSON.stringify(values))
                updatedData.sections = updatedData.sections.map((item, itemIndex) => {
                  if (itemIndex === sectionIndex) {
                    return { ...item, description: e.target.value }
                  }
                  return item
                })
                setValues(updatedData)
              }} />
            </div>
            <p>Foods <button onClick={() => {
              const newData = JSON.parse(JSON.stringify(values))
              newData.sections = newData.sections.map((x, i) => {
                if (i === sectionIndex) {
                  x.foods = [...x.foods, { food_id: null, price: null }]
                }
                return x
              })
              setValues(newData)
            }}>Add New Food</button></p>
            <div className="foodsInputWrapper">
              {section.foods.map((foodItem, foodItemIndex) => {
                return (<div className="foodInputs">
                  <div className="fomControll">
                    <Select
                      displayKey="name"
                      options={foodStore.list}
                      value={foodItem.food_id}
                      onChange={(value) => {
                        const updatedData = JSON.parse(JSON.stringify(values))
                        updatedData.sections = updatedData.sections.map((item, itemIndex) => {
                          if (itemIndex === sectionIndex) {
                            item.foods = item.foods.map((f, i) => {
                              if (i === foodItemIndex) {
                                return { ...f, food_id: value }
                              }
                              return f
                            })
                          }
                          return item
                        })
                        setValues(updatedData)
                      }}
                    />
                  </div>
                  <div className="fomControll">
                    <input type="text" id="price" value={foodItem.price}
                      onChange={(e) => {
                        e.preventDefault()
                        const updatedData = JSON.parse(JSON.stringify(values))
                        updatedData.sections = updatedData.sections.map((item, itemIndex) => {
                          if (itemIndex === sectionIndex) {
                            item.foods = item.foods.map((f, i) => {
                              if (i === foodItemIndex) {
                                return { ...f, price: e.target.value }
                              }
                              return f
                            })
                          }
                          return item
                        })
                        setValues(updatedData)
                      }} />
                  </div>

                </div>)
              })}
            </div>
          </section>
        })}
      </div>
    </form>
  )
}

export default MenuFormPage