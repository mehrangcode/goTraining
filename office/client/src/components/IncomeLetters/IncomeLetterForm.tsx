import IncomeLetterStore from "@src/pages/incomeLetters/store"
import { useEffect, useState } from "react"

const defaultValue = {
    number: 0,
    title: "",
    content: "",
    subjectId: "",
    created_At: "",
    owner: "",
    destination: "",
    status: null,
    operatorId: "",
}
function IncomeLetterForm() {
    const [values, setValues] = useState(JSON.parse(JSON.stringify(defaultValue)))
    const store = IncomeLetterStore()
    useEffect(() => {
        setValues({
            ...(store.targetItem || JSON.parse(JSON.stringify(defaultValue))),
        })
    }, [store.targetItem])
    async function onSubmitHandler(e) {
        e.preventDefault()
        try {
            if (store.targetItem?.id) {
                await store.update(store.targetItem.id, {
                    ...values,
                    number: +values.number
                })
            } else {
                await store.create({ ...values, number: +values.number })
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
        <form className="incomeLetterForm" onSubmit={onSubmitHandler}>
            <div className="row">
                <div className="col">
                    <div className="fomControll">
                        <label htmlFor="number">number</label>
                        <input type="text" id="number" value={values.number} onChange={(e) => {
                            e.preventDefault()
                            onChangeHandler("number", e.target.value)
                        }} />
                    </div>
                    <div className="fomControll">
                        <label htmlFor="title">title</label>
                        <input type="text" id="title" value={values.title} onChange={(e) => {
                            e.preventDefault()
                            onChangeHandler("title", e.target.value)
                        }} />
                    </div>
                    <div className="fomControll">
                        <label htmlFor="subjectId">subjectId</label>
                        <input type="text" id="subjectId" value={values.subjectId} onChange={(e) => {
                            e.preventDefault()
                            onChangeHandler("subjectId", e.target.value)
                        }} />
                    </div>
                </div>

                <div className="col">
                    <div className="fomControll">
                        <label htmlFor="subjectId">content</label>
                        <textarea id="content" value={values.content} onChange={(e) => {
                            e.preventDefault()
                            onChangeHandler("content", e.target.value)
                        }} />
                    </div>
                </div>
            </div>
            <button type="submit">{store.targetItem ? "Edit" : "Save"}</button>
            {store.targetItem ? <button type="button" onClick={() => store.selectUser(undefined)}>Reset</button> : null}
        </form>
    )
}

export default IncomeLetterForm