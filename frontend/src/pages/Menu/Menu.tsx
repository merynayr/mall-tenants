import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Menu.module.css';

const Menu: React.FC = () => {
	return (
		<div className={styles['menu-page']}>
			<h1>Меню</h1>
			<ul className="menu-list">
				<li><Link to='/premises'>Помещения</Link></li>
				<li><Link to="/clients">Арендаторы</Link></li>
				<li><Link to="/payments">Платежи</Link></li>
			</ul>
		</div>
	);
};

export default Menu;
