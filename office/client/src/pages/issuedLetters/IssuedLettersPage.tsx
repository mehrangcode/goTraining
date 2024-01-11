import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import IssuedStore from "./store"
import DoDate from "@doolooper/dodate"

function IssuedLettersPage() {
  const issuedStore = IssuedStore()
  useEffect(() => {
    issuedStore.fetchList()
  }, [])
  const nav = useNavigate()
  return (
    <div className="userPage">
      <button onClick={() => nav("modify")}>Create</button>
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
            <th>Subject</th>
            <th>Operator</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>

          {issuedStore.list.map(record => {
            return <tr>
              <td>{record.number}</td>
              <td>{record.title}</td>
              <td>{record.owner}</td>
              <td>{record.destination}</td>
              <td>{record.content}</td>
              <td>{record.status}</td>
              <td>{DoDate.parse(new Date(record.created_at).toISOString()).formatJalali("YYYY/MM/DD")}</td>
              <td>{record.subjectName}</td>
              <td>{record.operatorName}</td>
              <td>
                <div className="action">
                  <button onClick={() =>  nav("modify/"+record.id)}>Edit</button>
                  <button onClick={() => issuedStore.delete(record.id)}>Delete</button>
                </div>
              </td>
            </tr>
          })}
        </tbody>
      </table>
    </div>
  )
}

export default IssuedLettersPage