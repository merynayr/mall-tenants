import React from 'react';
import { NavLink } from 'react-router-dom';
import styles from './Premises.module.css';

const mockData: Record<number, { name: string; details: string }> = {
	1: { name: 'ООО "Ромашка"', details: 'Торговая точка №12, 1 этаж' },
	2: { name: 'ИП Иванов', details: 'Бутик 8, 2 этаж' }
};

const Premises: React.FC = () => {
	return (
		<>
			<div className={styles['page premises-list']}>
				<h2>Список помещений</h2>
				<ul>
					{Object.keys(mockData).map(id => {
						const premise = mockData[Number(id)];
						return (
							<li key={id}>
								<NavLink
									to={`/premises/${id}`}
									className={({ isActive }) => isActive ? styles.active : ''}
								>
									{premise.name}
								</NavLink>
							</li>
						);
					})}
				</ul>
			</div></>
	);
};

export default Premises;
