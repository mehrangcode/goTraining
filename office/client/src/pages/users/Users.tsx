import { useEffect } from "react"
import UserStore from "./store"

function Users() {
  const userStore = UserStore()
  useEffect(() => {
    userStore.fetchList()
  }, [])
  return (
    <div className="userPage">
      <table className="list">
        <thead>
          <tr>
            <th>Name</th>
            <th>email</th>
            <th>action</th>
          </tr>
        </thead>
        <tbody>

          {userStore.list.map(record => {
            return <tr>
              <td>{record.name}</td>
              <td>{record.email}</td>
              <td><button>Delete</button></td>
            </tr>
          })}
        </tbody>
      </table>
    </div>
  )
}

export default Users