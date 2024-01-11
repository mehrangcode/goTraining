import { useEffect } from "react"
import SubjectStore from "./store"
import SubjectForm from "@src/components/subjects/SubjectForm"

function SubjectsPage() {
  const subjectStore = SubjectStore()
  useEffect(() => {
    subjectStore.fetchList()
  }, [])
  return (
    <div className="subjectPage">
      <SubjectForm />
      <table className="list">
        <thead>
          <tr>
            <th>Label</th>
            <th>Archive</th>
            <th>action</th>
          </tr>
        </thead>
        <tbody>

          {subjectStore.list?.map(record => {
            return <tr>
              <td>{record.label}</td>
              <td>{record.archive}</td>
              <td>
                <div className="action">
                  <button onClick={() => subjectStore.selectSubject(record)}>Edit</button>
                  <button onClick={() => subjectStore.delete(record.id)}>Delete</button>
                </div>
              </td>
            </tr>
          })}
        </tbody>
      </table>
    </div>
  )
}

export default SubjectsPage