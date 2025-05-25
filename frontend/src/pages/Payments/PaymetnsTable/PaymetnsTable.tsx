import React from 'react';
import styles from './PaymetnsTable.module.css';
import { Payment } from '@/interfaces/payment';

interface Props {
	payments: Payment[];
}

export const PaymentsTable: React.FC<Props> = ({ payments }) => {

	function getPaidUntilDate(startDate: string, paidMonths: number): string {
		const date = new Date(startDate);
		date.setMonth(date.getMonth() + paidMonths);
		return date.toLocaleDateString();
	}

	
	return (
			<table className={styles.table}>
			<thead>
				<tr>
					<th>ID</th>
					<th>Договор</th>
					<th>Объект аренды</th>
					<th>Клиент</th>
					<th>Период</th>
					<th>Сумма</th>
					<th>Статус</th>
					<th>Оплата с</th>
					<th>Оплата с</th>
					<th>Дата</th>
				</tr>
			</thead>
			<tbody>
				{payments.map((payment) => (
					<tr key={payment.id}>
						<td>{payment.id}</td>
						<td>{payment.rental_id}</td>
						<td>{payment.space_code}</td>
						<td>{payment.client_name}</td>
						<td>{getPaidUntilDate(payment.start_date, payment.paid_months)}</td>
						<td>{payment.amount}</td>
						<td>{payment.status}</td>
						<td>{new Date(payment.start_date).toLocaleDateString()}</td>
						<td>{new Date(payment.end_date).toLocaleDateString()}</td>
						<td>{new Date(payment.created_at).toLocaleDateString()}</td>
					</tr>
				))}
			</tbody>
		</table>
	);
};
