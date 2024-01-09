import SubjectStore from "@src/pages/subjects/store"
import { useEffect, useState } from "react"

function SubjectForm() {
    const [values, setValues] = useState({
        label: ""
    })
    const subjectStore = SubjectStore()
    useEffect(() => {
        setValues({
            ...(subjectStore.targetItem || {
                label: "",
                archive: false
            }),
        })
    }, [subjectStore.targetItem])
    async function onSubmitHandler(e) {
        e.preventDefault()
        try {
            const payload = {
                label: values.label
            }
            if (subjectStore.targetItem?.id) {
                await subjectStore.update(subjectStore.targetItem.id, {
                    label: values.label
                })
            } else {
                await subjectStore.create(payload)
            }
            setValues({
                label: ""
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
        <form className="subjectForm" onSubmit={onSubmitHandler}>
            <div className="fomControll">
                <label htmlFor="label">Subject</label>
                <input type="text" id="label" value={values.label} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("label", e.target.value)
                }} />
            </div>
            <button type="submit">{subjectStore.targetItem ? "Edit" : "Save"}</button>
            {subjectStore.targetItem ? <button type="button" onClick={() => subjectStore.selectSubject(undefined)}>Reset</button> : null}
        </form>
    )
}

export default SubjectForm