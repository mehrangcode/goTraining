import { Link } from "react-router-dom"

function Home() {
  return (
    <div className="homePage">
        <h1>Home</h1>
        <h3><Link to={"/users"}>USERS</Link></h3>
        <hr />
        <h1>letters</h1>
        <h3><Link to={"/letters/income"}>income</Link></h3>
    </div>
  )
}

export default Home