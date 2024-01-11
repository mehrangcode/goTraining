import IssuedLetterStore from "@src/pages/issuedLetters/store"
import { useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"
import Select from "../shared/select/Select"

const defaultValue = {
    number: undefined,
    title: "",
    content: "",
    subjectId: "",
    created_At: "",
    owner: "",
    destination: "",
    status: null,
    operatorId: "",
    template: ""
}
function IssuedLetterForm() {
    const [values, setValues] = useState(JSON.parse(JSON.stringify(defaultValue)))
    const store = IssuedLetterStore()
    const { letterId } = useParams()
    console.log(letterId)
    useEffect(() => {
        if (letterId) {
            store.fetchItemById(letterId)
        }
    }, [letterId])
    useEffect(() => {
        setValues({
            ...(store.targetItem || JSON.parse(JSON.stringify(defaultValue))),
        })
    }, [store.targetItem])

    const nav = useNavigate()
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
            nav("/letters/issued")
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
        <form className="issuedLetterForm" onSubmit={onSubmitHandler} autoComplete="off">
            <div className="sidebar">
                <div className="formControll">
                    <label htmlFor="template">Template</label>
                    <Select
                        value={values.temp}
                        options={[
                            { id: "temp1", label: "Temp 1" },
                            { id: "temp2", label: "Temp 2" },
                            { id: "temp3", label: "Temp 3" },
                            { id: "temp4", label: "Temp 4" }
                        ]}
                        onChange={(value: string) => {
                            onChangeHandler("temp", value)
                        }}
                    />
                </div>
                <div className="fomControll">
                    <label htmlFor="number">number</label>
                    <input type="text" id="number" value={values.number} onChange={(e) => {
                        e.preventDefault()
                        onChangeHandler("number", e.target.value)
                    }} />
                </div>
                <div className="fomControll">
                    <label htmlFor="issuedDate">issuedDate</label>
                    <input type="text" id="issuedDate" value={values.issuedDate} onChange={(e) => {
                        e.preventDefault()
                        onChangeHandler("issuedDate", e.target.value)
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
                <div className="actionFooter">
                    <button type="submit">{store.targetItem ? "Edit" : "Save"}</button>
                    <button type="button" onClick={() => {
                        store.selectUser(undefined)
                        nav("/letters/issued")
                    }}>Cancel</button>
                </div>
            </div>
            <textarea id="content" value={values.content} onChange={(e) => {
                e.preventDefault()
                onChangeHandler("content", e.target.value)
            }} />

        </form>
    )
}

export default IssuedLetterForm