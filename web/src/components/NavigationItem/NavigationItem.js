import { Link } from 'react-router-dom';
import './NavigationItem.css'

const NavigationItem = ({ name, url }) => {
    return (
        <li className='navigation__item'>
          <Link to={url}>{name}</Link>
        </li>
    )
}

export default NavigationItem;