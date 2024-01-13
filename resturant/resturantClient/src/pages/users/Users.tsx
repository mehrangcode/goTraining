import { useEffect } from "react"
import UserStore from "./store"
import UserForm from "@src/components/users/UserForm"

function Users() {
  const userStore = UserStore()
  useEffect(() => {
    userStore.fetchList()
  }, [])
  return (
    <div className="userPage">
      <UserForm />
      <table className="list">
        <thead>
          <tr>
            <th>Name</th>
            <th>email</th>
            <th>action</th>
          </tr>
        </thead>
        <tbody>

          {userStore.list?.map(record => {
            return <tr>
              <td>{record.name}</td>
              <td>{record.email}</td>
              <td>
                <div className="action">
                  <button onClick={() => userStore.selectUser(record)}>Edit</button>
                  <button onClick={() => userStore.delete(record.id)}>Delete</button>
                </div>
              </td>
            </tr>
          })}
        </tbody>
      </table>
    </div>
  )
}

export default Users