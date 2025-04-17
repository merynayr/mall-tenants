import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Clients.module.css';

const Home: React.FC = () => {
	return (
		<div className={styles['page home']}>
			<h1>Добро пожаловать</h1>
			<p>Выберите раздел:</p>
			<ul>
				<li><Link to="/tenants">Арендаторы</Link></li>
			</ul>
		</div>
	);
};

export default Home;
