import TableStore, { TableType } from "@src/pages/tables/store"
import { useEffect, useState } from "react"
const defaultValue: TableType = {
    name: "",
    capacity: 0,
    status: 1,
    photos: []
}
function TableForm() {
    const [values, setValues] = useState<TableType>(JSON.parse(JSON.stringify(defaultValue)))
    const tableStore = TableStore()
    useEffect(() => {
        setValues({
            ...(tableStore.targetItem || JSON.parse(JSON.stringify(defaultValue))),
        })
    }, [tableStore.targetItem])
    async function onSubmitHandler(e) {
        e.preventDefault()
        try {
            const payload = {
                name: values.name,
                capacity: +values.capacity,
                status: values.status,
                photos: values.photos
            }
            if (tableStore.targetItem?.id) {
                payload.status = undefined
                await tableStore.update(tableStore.targetItem.id, payload)
            } else {
                await tableStore.create(payload)
            }
            setValues(JSON.parse(JSON.stringify(defaultValue)))
        } catch (error) {
        }
    }

    function onChangeHandler(name: string, value: string | string[]) {
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
                <label htmlFor="capacity">Capacity</label>
                <input type="text" id="capacity" value={values.capacity} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("capacity", e.target.value)
                }} />
            </div>
            <div className="fomControll">
                <label htmlFor="photos">Photos</label>
                <input type="text" id="photos" value={values.photos} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("photos", e.target.value)
                }} />
            </div>
            {!tableStore.targetItem ? <div className="fomControll">
                <label htmlFor="status">status</label>
                <input type="text" id="status" value={values.status} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("status", e.target.value)
                }} />
            </div> : null}
            <button type="submit">{tableStore.targetItem ? "Edit" : "Save"}</button>
            {tableStore.targetItem ? <button type="button" onClick={() => tableStore.selectTable(undefined)}>Reset</button> : null}
        </form>
    )
}

export default TableForm