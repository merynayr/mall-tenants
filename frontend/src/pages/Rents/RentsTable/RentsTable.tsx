import React from 'react';
import styles from './RentsTable.module.css';
import PremiseCard from '@/components/PremisesCard/PremisesCard';
import { Rent } from '@/interfaces/rent';
import api from '@/helpers/API';

interface Props {
	rents: Rent[];
}

export const RentsTable: React.FC<Props> = ({ rents }) => {
	function getPaidUntilDate(startDate: string, paidMonths: number): string {
		const date = new Date(startDate);
		date.setMonth(date.getMonth() + paidMonths);
		return date.toLocaleDateString();
	}

	const handleDownload = async (contractId: number) => {
		try {
			const response = await api.get(`/contracts/${contractId}`, {
				responseType: 'blob',
			});

			const file = new Blob([response.data], { type: 'application/pdf' });
			const url = window.URL.createObjectURL(file);

			const link = document.createElement('a');
			link.href = url;
			link.download = `contract-${contractId}.pdf`;
			document.body.appendChild(link);
			link.click();
			link.remove();

			window.URL.revokeObjectURL(url);
		} catch (error) {
			console.error('Ошибка при загрузке договора:', error);
		}
	};

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
					<th>Договоры</th>
				</tr>
			</thead>
			<tbody>
				{rents.map((rent) => (
					<tr key={rent.rental_id}>
						<td>{rent.rental_id}</td>
						<td>{rent.client_name}</td>
						<td>
							<PremiseCard
								code={rent.space_code}
								status={'available'}
								floor={undefined}
								area={undefined}
							/>
						</td>
						<td>{new Date(rent.created_at).toLocaleDateString()}</td>
						<td>{new Date(rent.start_date).toLocaleDateString()}</td>
						<td>{new Date(rent.end_date).toLocaleDateString()}</td>
						<td>{getPaidUntilDate(rent.start_date, rent.paid_months)}</td>
						<td>{rent.paid_months} мес</td>
						<td>
							{rent.contracts && rent.contracts.length > 0 ? (
								rent.contracts.map((contract) => (
									<button
										key={contract.contract_id}
										onClick={() => handleDownload(contract.contract_id)}
										title="Скачать договор"
										className={styles.iconButton}
									>
										📄
									</button>
								))
							) : (
								<span>—</span>
							)}
						</td>
					</tr>
				))}
			</tbody>
		</table>
	);
};
