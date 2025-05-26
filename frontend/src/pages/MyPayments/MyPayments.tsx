import { useEffect, useState } from 'react';
import styles from './MyPayments.module.css';
import Button from '@/components/Button/Button';
import api from '@/helpers/API';
import { Payment } from '@/interfaces/payment';
import { PaymentsTable } from '@/pages/Payments/PaymetnsTable/PaymetnsTable';
import { AxiosError } from 'axios';
import { RootState } from '@/store/store';
import { useSelector } from 'react-redux';
import { toast } from 'react-toastify';

export function PageMyPayments() {
	const [payments, setPayments] = useState<Payment[]>([]);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const [limit] = useState(20);
	const [offset, setOffset] = useState(0);
	const [status, setStatus] = useState<'all' | 'paid' | 'unpaid'>('all');
	const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('desc');
	const { profile } = useSelector((state: RootState) => state.user);

	useEffect(() => {
		if (profile?.user_id) {
			fetchPayments(profile.user_id);
		}
		setError("");
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [offset, status, sortOrder, profile?.user_id]);

	useEffect(() => {
		if (error) {
			toast.error(error);
		}
	}, [error]);
	
	const fetchPayments = async (id: number | undefined) => {
		try {
			setIsLoading(true);
			const isPaidParam = status === 'all' ? undefined : status === 'paid';

			const { data } = await api.get<Payment[]>(`/payments/client/${id}`, {
				params: {
					limit,
					offset,
					is_paid: isPaidParam,
					sort_by: 'period_start',
					sort_order: sortOrder,
				}
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

		<div className={styles.filters}>
			<select value={status} onChange={(e) => setStatus(e.target.value as any)}>
				<option value="all">Все</option>
				<option value="paid">Оплаченные</option>
				<option value="unpaid">Неоплаченные</option>
			</select>

			<select value={sortOrder} onChange={(e) => setSortOrder(e.target.value as any)}>
				<option value="desc">Сначала новые</option>
				<option value="asc">Сначала старые</option>
			</select>
		</div>

		<div>
			{isLoading ? (
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

export default PageMyPayments;
