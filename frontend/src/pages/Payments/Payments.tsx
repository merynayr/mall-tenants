import { useEffect, useState } from 'react';
import styles from './Payments.module.css';
import api from '@/helpers/API';
import { Payment } from '@/interfaces/payment';
import { PaymentsTable } from '@/pages/Payments/PaymetnsTable/PaymetnsTable';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import Pagination from '@/components/Pagination/Pagination';

export function PagePayments() {
	const [payments, setPayments] = useState<Payment[]>([]);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const [limit] = useState(11);
	const [offset, setOffset] = useState(0);
	const [status, setStatus] = useState<'all' | 'paid' | 'unpaid'>('all');
	const [searchQuery, setSearchQuery] = useState('');
	const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('asc');
	const [showOverdue, setShowOverdue] = useState(false);

	useEffect(() => {
		fetchPayments();
		setError("");
	// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [offset, status, searchQuery, sortOrder, showOverdue]);

	useEffect(() => {
		if (error) {
			toast.error(error);
		}
	}, [error]);
	
	const fetchPayments = async () => {
		try {
			setIsLoading(true);
			const isPaidParam = status === 'all' ? undefined : status === 'paid';
			const { data } = await api.get<Payment[]>('/payments', {
				params: {
					limit,
					offset,
					is_paid: isPaidParam,
					client: searchQuery.trim() || undefined,
					sort_by: 'period_start',
					sort_order: sortOrder,
					is_overdue: showOverdue || undefined,
				}
			});

			setPayments(data);
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.response?.data.error);
				if (e.response?.data.error === "Платежи не найдены") {
      		setPayments([]);
				}
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
			<input
				type="text"
				placeholder="Поиск по имени клиента..."
				value={searchQuery}
				onChange={(e) => setSearchQuery(e.target.value)}
			/>

			<select value={status} onChange={(e) => setStatus(e.target.value as any)}>
				<option value="all">Все</option>
				<option value="paid">Оплаченные</option>
				<option value="unpaid">Неоплаченные</option>
			</select>

			<select value={sortOrder} onChange={(e) => setSortOrder(e.target.value as any)}>
				<option value="desc">Сначала новые</option>
				<option value="asc">Сначала старые</option>
			</select>

			<div className={styles.checkboxWrapper}>
				<label className={styles.checkboxLabel}>
					<input
						type="checkbox"
						checked={showOverdue}
						onChange={(e) => setShowOverdue(e.target.checked)}
					/>
					<span>Просроченные</span>
				</label>
			</div>
		</div>

		<div>
			{isLoading ? (
				<div>Загрузка...</div>
			) : (
				<>
					
					<PaymentsTable payments={payments} />
					
					<Pagination
						offset={offset}
						limit={limit}
						total={payments.length}
						onNext={handleNext}
						onPrev={handlePrev}
					/>
				</>
			)}
		
		</div>
	</>;
}

export default PagePayments;
