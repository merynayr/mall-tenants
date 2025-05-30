import React from 'react';
import styles from './Pagination.module.css';
import Button from '@/components/Button/Button';

interface PaginationProps {
	offset: number;
	limit: number;
	total: number;
	onNext: () => void;
	onPrev: () => void;
}

const Pagination: React.FC<PaginationProps> = ({ offset, limit, total, onNext, onPrev }) => {
	const from = offset + 1;
	const to = offset + total;

	return (
		<div className={styles.pagination}>
   		<Button className={styles.small} onClick={onPrev} disabled={offset === 0}>Назад</Button>
      <span>Показано с {from} по {to}</span>
			<Button className={styles.small} onClick={onNext} disabled={total < limit}>Вперёд</Button>
		</div>
	);
};

export default Pagination;
