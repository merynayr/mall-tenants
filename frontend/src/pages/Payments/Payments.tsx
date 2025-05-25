import { useEffect, useState } from 'react';
import styles from './Payments.module.css';
import Button from '@/components/Button/Button';
import api from '@/helpers/API';
import { Payment } from '@/interfaces/payment';
import { PaymentsTable } from '@/pages/Payments/PaymetnsTable/PaymetnsTable';
import { AxiosError } from 'axios';

export function PagePayments() {
	const [payments, setPayments] = useState<Payment[]>([]);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const [limit] = useState(20);
	const [offset, setOffset] = useState(0);
	
	useEffect(() => {
		fetchPayments();
	// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [offset]);

	const fetchPayments = async () => {
		try {
			setIsLoading(true);
			const { data } = await api.get<Payment[]>('/payments', {
				params: { limit, offset }
			});
			setPayments(data);
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.response?.data.error);
			}		
		} finally {
			setIsLoading(false);
		}
	};


	const handleNext = () => setOffset((prev) => prev + limit);
	const handlePrev = () => setOffset((prev) => Math.max(prev - limit, 0));

	return <>
	
		<div className={styles.headRow}>
			<h1>Список платёжных операций</h1>
		</div>

		<div>
			{error ? <div className={styles.error}>{error}</div> :
			isLoading ? (
				<div>Загрузка...</div>
			) : (
				<>
					<PaymentsTable payments={payments} />
					<div className={styles.pagination}>
						<Button className={styles['small']} onClick={handlePrev} disabled={offset === 0}>Назад</Button>
						<span>Показано с {offset + 1} по {offset + payments.length}</span>
						<Button className={styles['small']} onClick={handleNext} disabled={payments.length < limit}>Вперёд</Button>
					</div>
				</>
			)}
		
		</div>
	</>;
}

export default PagePayments;
