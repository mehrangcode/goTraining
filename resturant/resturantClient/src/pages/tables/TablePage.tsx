import { useEffect } from "react"
import TableStore from "./store"
import TableForm from "./components/TableForm"

function TablesPage() {
  const tableStore = TableStore()
  useEffect(() => {
    tableStore.fetchList()
  }, [])
  return (
    <div className="tablePage">
      <TableForm />
      <table className="list">
        <thead>
          <tr>
            <th>Name</th>
            <th>Capacity</th>
            <th>Status</th>
            <th>action</th>
          </tr>
        </thead>
        <tbody>

          {tableStore.list?.map(record => {
            return <tr>
              <td>{record.name}</td>
              <td>{record.capacity}</td>
              <td>{record.status}</td>
              <td>
                <div className="action">
                  <button onClick={() => tableStore.selectTable(record)}>Edit</button>
                  <button onClick={() => tableStore.delete(record.id)}>Delete</button>
                </div>
              </td>
            </tr>
          })}
        </tbody>
      </table>
    </div>
  )
}

export default TablesPage