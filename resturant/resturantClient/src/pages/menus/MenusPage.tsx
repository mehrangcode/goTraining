import { useEffect } from "react"
import MenuStore from "./store"
import MenuForm from "./components/MenuForm"

function MenusPage() {
  const menuStore = MenuStore()
  useEffect(() => {
    menuStore.fetchList()
  }, [])
  return (
    <div className="menuPage">
      <MenuForm />
      <table className="list">
        <thead>
          <tr>
            <th>Name</th>
            <th>email</th>
            <th>status</th>
            <th>action</th>
          </tr>
        </thead>
        <tbody>

          {menuStore.list?.map(record => {
            return <tr>
              <td>{record.title}</td>
              <td>{record.description}</td>
              <td>{record.status}</td>
              <td>
                <div className="action">
                  <button onClick={() => menuStore.selectMenu(record)}>Edit</button>
                  <button onClick={() => menuStore.delete(record.id)}>Delete</button>
                </div>
              </td>
            </tr>
          })}
        </tbody>
      </table>
    </div>
  )
}

export default MenusPage