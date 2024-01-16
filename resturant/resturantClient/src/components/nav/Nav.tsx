import { Link } from 'react-router-dom'

function Nav() {
    return (
        <nav className="nav"><Link to={"./"}>OFFICE</Link>
        <ul className="topMenu">
            <li><Link to="/users">Users</Link> </li>
            <li><Link to="/foods">foods</Link> </li>
            <li><Link to="/menus">menus</Link> </li>
        </ul>
        </nav>
    )
}

export default Nav