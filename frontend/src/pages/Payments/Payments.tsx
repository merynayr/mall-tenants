import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Payments.module.css';

type Tenant = {
  id: number;
  name: string;
  contractEnd: string;
};

const tenants: Tenant[] = [
	{ id: 1, name: 'ООО "Ромашка"', contractEnd: '2025-12-01' },
	{ id: 2, name: 'ИП Иванов', contractEnd: '2024-09-15' }
];

const Tenants: React.FC = () => {
	return (
		<div className={styles['page tenants']}>
			<h2>Список оплат</h2>
			<ul>
				{tenants.map((tenant) => (
					<li key={tenant.id}>
						<Link to={`/tenants/${tenant.id}`}>
							{tenant.name} — до {tenant.contractEnd}
						</Link>
					</li>
				))}
			</ul>
		</div>
	);
};

export default Tenants;
