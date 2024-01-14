import UserStore from "@src/pages/users/store"
import { useEffect, useState } from "react"

function UserForm() {
    const [values, setValues] = useState({
        name: "",
        email: "",
        password: "",
    })
    const userStore = UserStore()
    useEffect(() => {
        setValues({
            ...(userStore.targetItem || {
                name: "",
                email: "",
                password: ""
            }),
        })
    }, [userStore.targetItem])
    async function onSubmitHandler(e) {
        e.preventDefault()
        try {
            const payload = {
                name: values.name,
                email: values.email,
                password: values.password
            }
            if (userStore.targetItem?.id) {
                await userStore.update(userStore.targetItem.id, {
                    name: values.name,
                    email: values.email,
                    password: undefined
                })
            } else {
                await userStore.create(payload)
            }
            setValues({
                name: "",
                email: "",
                password: ""
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
        <form className="horizontal" onSubmit={onSubmitHandler}>
            <div className="fomControll">
                <label htmlFor="name">Name</label>
                <input type="text" id="name" value={values.name} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("name", e.target.value)
                }} />
            </div>
            <div className="fomControll">
                <label htmlFor="email">Email</label>
                <input type="text" id="email" value={values.email} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("email", e.target.value)
                }} />
            </div>
            {!userStore.targetItem ? <div className="fomControll">
                <label htmlFor="password">Password</label>
                <input type="password" id="password" value={values.password} onChange={(e) => {
                    e.preventDefault()
                    onChangeHandler("password", e.target.value)
                }} />
            </div> : null}
            <button type="submit">{userStore.targetItem ? "Edit" : "Save"}</button>
            {userStore.targetItem ? <button type="button" onClick={() => userStore.selectUser(undefined)}>Reset</button> : null}
        </form>
    )
}

export default UserForm