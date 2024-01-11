import IssuedLetterStore from "@src/pages/issuedLetters/store"
import { useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"

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
            <div className="row letterFormHeaderSection">
                <div className="col numberSection">
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
                </div>
                <div className="col titleSection">
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
            </div>
            <div className="fomControll">
                <label htmlFor="subjectId">content</label>
                <textarea id="content" value={values.content} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("content", e.target.value)
                }} />
            </div>
            <div className="actionFooter">
                <button type="submit">{store.targetItem ? "Edit" : "Save"}</button>
                <button type="button" onClick={() => {
                    store.selectUser(undefined)
                    nav("/letters/issued")
                }}>Cancel</button>
            </div>
        </form>
    )
}

export default IssuedLetterForm