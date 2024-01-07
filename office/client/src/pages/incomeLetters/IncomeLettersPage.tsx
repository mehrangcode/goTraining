import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import UserStore from "./store"

function IncomeLettersPage() {
  const incomeStore = UserStore()
  useEffect(() => {
    incomeStore.fetchList()
  }, [])
  const nav = useNavigate()
  return (
    <div className="userPage">
      <button onClick={() => nav("/income/modify")}>Create</button>
      <table className="list">
        <thead>
          <tr>
            <th>Number</th>
            <th>Title</th>
            <th>Owner</th>
            <th>Destination</th>
            <th>Content</th>
            <th>status</th>
            <th>Created Date</th>
            <th>SubjectId</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>

          {incomeStore.list.map(record => {
            return <tr>
              <td>{record.number}</td>
              <td>{record.title}</td>
              <td>{record.owner}</td>
              <td>{record.destination}</td>
              <td>{record.content}</td>
              <td>{record.status}</td>
              <td>{record.created_At}</td>
              <td>{record.subjectId}</td>
              <td>
                <div className="action">
                  <button onClick={() => incomeStore.selectUser(record)}>Edit</button>
                  <button onClick={() => incomeStore.delete(record.id)}>Delete</button>
                </div>
              </td>
            </tr>
          })}
        </tbody>
      </table>
    </div>
  )
}

export default IncomeLettersPage