import React from 'react';
import styles from './PaymetnsTable.module.css';
import { Payment } from '@/interfaces/payment';
import { calculateRentalDuration } from '@/helpers/CountMonth';

interface Props {
	payments: Payment[];
}

export const PaymentsTable: React.FC<Props> = ({ payments }) => {
	return (
			<table className={styles.table}>
			<thead>
				<tr>
					<th>ID</th>
					<th>Договор</th>
					<th>Объект аренды</th>
					<th>Клиент</th>
					<th>Период</th>
					<th>Сумма ₽</th>
					<th>Статус</th>
					<th>Оплата с</th>
					<th>Оплата до</th>
					<th>Дата оплаты</th>
				</tr>
			</thead>
			<tbody>
				{payments.map((payment) => (
					<tr key={payment.id}>
						<td>{payment.id}</td>
						<td>{payment.rental_id}</td>
						<td>{payment.space_code}</td>
						<td>{payment.client_name}</td>
						<td>{calculateRentalDuration(payment.start_date, payment.end_date)} мес.</td>
						<td>{payment.amount}</td>
						<td>{payment.status ? 'Оплачено' : 'Не оплачено'}</td>
						<td>{new Date(payment.start_date).toLocaleDateString()}</td>
						<td>{new Date(payment.end_date).toLocaleDateString()}</td>
						<td>{payment.payment_date != "" ? "" : new Date(payment.payment_date).toLocaleDateString()}</td>
					</tr>
				))}
			</tbody>
		</table>
	);
};
