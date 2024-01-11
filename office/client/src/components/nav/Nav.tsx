import { Link } from 'react-router-dom'

function Nav() {
    return (
        <nav className="nav"><Link to={"./"}>OFFICE</Link>
        <ul className="topMenu">
            <li><Link to="/users">Users</Link> </li>
            <li><Link to="/Subjects">Subjects</Link> </li>
            <li><Link to="/letters/Income">Income Letters</Link> </li>
            <li><Link to="/letters/Issued">Issued Letters</Link> </li>
        </ul>
        </nav>
    )
}

export default Nav