// import React from 'react';
// import { useNavigate } from 'react-router-dom';
// import { Rent } from '@/interfaces/rent';
// import styles from './RentCard.module.css';
// import { formatDate } from '@/helpers/FormatDate';


// interface RentCardProps {
// 	rent: Rent;
// }

// const RentCard: React.FC<RentCardProps> = ({ rent }) => {
// 	const navigate = useNavigate();
	
// 	const handlePayClick = () => {
// 		navigate(`/my-payments?rental_id=${rent.rental_id}&status=unpaid&sort=asc`);
// 	};

// 	return (
// 		<div className={styles['rent-card']}>
// 			<div className={styles['rent-title']}>
// 				<span role="img" aria-label="building">🏢</span>
// 				Помещение: {rent.space_code}
// 			</div>

// 			<p className={styles.period}>
// 				<span role="img" aria-label="calendar">📅</span>&nbsp;
// 				<span className={styles.date}>{formatDate(rent.start_date)}</span>
// 				&nbsp;—&nbsp;
// 				<span className={styles.date}>{formatDate(rent.end_date)}</span>
// 			</p>

// 			<div className={styles['rent-footer']}>
// 				<span className={styles.badge}>Оплачено: {rent.paid_months} мес.</span>
// 				<span className={styles.badge}>Создано: {formatDate(rent.created_at)}</span>
// 				{rent.updated_at && (
// 					<span className={styles.badge}>Обновлена: {formatDate(rent.updated_at)}</span>
// 				)}
// 			</div>

// 			<div className={styles['rent-actions']}>
// 				<button className={styles['action-btn']} onClick={handlePayClick}>
// 					<span className={styles['action-icon']}>💳</span>
// 					<span>Оплатить</span>
// 				</button>
// 				<button className={styles['action-btn']}>
// 					<span className={styles['action-icon']}>🔁</span>
// 					<span>Продлить</span>
// 				</button>
// 				<button className={styles['action-btn']}>
// 					<span className={styles['action-icon']}>❌</span>
// 					<span>Отменить</span>
// 				</button>
// 			</div>

// 		</div>
// 	);
// };

// export default RentCard;

import React from 'react';
import { useNavigate } from 'react-router-dom';
import { RentalWithContract } from '@/interfaces/rent';
import styles from './RentCard.module.css';
import { formatDate } from '@/helpers/FormatDate';
import api from '@/helpers/API';

interface RentCardProps {
	rent: RentalWithContract;
}

const RentCard: React.FC<RentCardProps> = ({ rent }) => {
	const navigate = useNavigate();

	const handlePayClick = () => {
		navigate(`/my-payments?rental_id=${rent.rental_id}&status=unpaid&sort=asc`);
	};

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

			// очистка объекта URL после загрузки
			window.URL.revokeObjectURL(url);
		} catch (error) {
			console.error('Ошибка при загрузке договора:', error);
		}
	};


	const handleSign = (contractId: number) => {
		// TODO: открыть модалку или отправить запрос на подписание
		alert(`Подписание договора №${contractId}`);
	};

	return (
		<div className={styles['rent-card']}>
			<div className={styles['rent-title']}>
				<span role="img" aria-label="building">🏢</span>
				Помещение: {rent.space_code}
			</div>

			<p className={styles.period}>
				<span role="img" aria-label="calendar">📅</span>&nbsp;
				<span className={styles.date}>{formatDate(rent.start_date)}</span>
				&nbsp;—&nbsp;
				<span className={styles.date}>{formatDate(rent.end_date)}</span>
			</p>

			<div className={styles['rent-footer']}>
				<span className={styles.badge}>Создано: {formatDate(rent.created_at)}</span>
			</div>

			<div className={styles['contract-section']}>
				<h4>Договоры</h4>
				{rent.contracts?.length === 0 && <p>Нет договоров</p>}
				{rent.contracts?.map(contract => (
					<div key={contract.contract_id} className={styles['contract-item']}>
						<span>ID: {contract.contract_id}</span>
						<button
							className={styles['action-btn']}
							onClick={() => handleDownload(contract.contract_id)}
						>
							📄 Скачать
						</button>
						{!contract.is_signed && (
							<button
								className={styles['action-btn']}
								onClick={() => handleSign(contract.contract_id)}
							>
								✍️ Подписать
							</button>
						)}
					</div>
				))}
			</div>

			<div className={styles['rent-actions']}>
				<button className={styles['action-btn']} onClick={handlePayClick}>
					💳 Оплатить
				</button>
				<button className={styles['action-btn']}>🔁 Продлить</button>
				<button className={styles['action-btn']}>❌ Отменить</button>
			</div>
		</div>
	);
};

export default RentCard;
