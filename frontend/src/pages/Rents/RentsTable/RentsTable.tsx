import React from 'react';
import styles from './RentsTable.module.css';
import PremiseCard from '@/components/PremisesCard/PremisesCard';
import { Rent } from '@/interfaces/rent';

interface Props {
	rents: Rent[];
}

export const RentsTable: React.FC<Props> = ({ rents }) => {

	function getPaidUntilDate(startDate: string, paidMonths: number): string {
		const date = new Date(startDate);
		date.setMonth(date.getMonth() + paidMonths);
		return date.toLocaleDateString();
	}

	
	return (
		<table className={styles.table}>
			<thead>
				<tr>
					<th>№ договора</th>
					<th>Клиент</th>
					<th>Помешение</th>
					<th>Дата заключения</th>
					<th>Начало аренды</th>
					<th>Конец аренды</th>
					<th>Оплачено до</th>
					<th>Период</th>
				</tr>
			</thead>
			<tbody>
				{rents.map((rent) => (
					<tr key={rent.rental_id}>
						<td>{rent.rental_id}</td>
						<td>{rent.client_name}</td>
						<td>
							<PremiseCard
								code={rent.space_code} status={'available'} floor={undefined} area={undefined}	
							/>
						</td>
						<td>{new Date(rent.created_at).toLocaleDateString()}</td>
						<td>{new Date(rent.start_date).toLocaleDateString()}</td>
						<td>{new Date(rent.end_date).toLocaleDateString()}</td>
						<td>{getPaidUntilDate(rent.start_date, rent.paid_months)}</td>
						<td>{rent.paid_months} мес</td>
					</tr>
				))}
			</tbody>
		</table>
	);
};
