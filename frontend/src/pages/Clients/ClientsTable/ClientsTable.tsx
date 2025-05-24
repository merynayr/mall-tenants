import React from 'react';
import styles from './ClientsTable.module.css';
import { Client } from '@/interfaces/client';

interface Props {
  clients: Client[];
}

export const ClientsTable: React.FC<Props> = ({ clients }) => {
	return (
		<table className={styles.table}>
			<thead>
				<tr>
					<th>ID</th>
					<th>Организация</th>
					<th>Контактное лицо</th>
					<th>Адрес</th>
					<th>Телефон</th>
					<th>Реквизиты</th>
				</tr>
			</thead>
			<tbody>
				{clients.map(client => (
					<tr key={client.id}>
						<td>{client.id}</td>
						<td>{client.organizationName}</td>
						<td>{client.contactPerson}</td>
						<td>{client.address}</td>
						<td>{client.phone}</td>
						<td>{client.requisites}</td>
					</tr>
				))}
			</tbody>
		</table>
	);
};
