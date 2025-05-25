import React from 'react';
import { Rent } from '@/interfaces/rent';
import styles from './RentCard.module.css';

interface RentCardProps {
	rent: Rent;
}

const RentCard: React.FC<RentCardProps> = ({ rent }) => {
	const formatDate = (dateString: string) => {
		const date = new Date(dateString);
		const day = String(date.getDate()).padStart(2, '0');
		const month = String(date.getMonth() + 1).padStart(2, '0');
		const year = String(date.getFullYear());
		return `${day}.${month}.${year}`;
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
				<span className={styles.badge}>Оплачено: {rent.paid_months} мес.</span>
				<span className={styles.badge}>Создано: {formatDate(rent.created_at)}</span>
				{rent.updated_at && (
					<span className={styles.badge}>Обновлена: {formatDate(rent.updated_at)}</span>
				)}
			</div>

			<div className={styles['rent-actions']}>
				<button className={styles['action-btn']}>
					<span className={styles['action-icon']}>💳</span>
					<span>Оплатить</span>
				</button>
				<button className={styles['action-btn']}>
					<span className={styles['action-icon']}>🔁</span>
					<span>Продлить</span>
				</button>
				<button className={styles['action-btn']}>
					<span className={styles['action-icon']}>❌</span>
					<span>Отменить</span>
				</button>
			</div>

		</div>
	);
};

export default RentCard;
