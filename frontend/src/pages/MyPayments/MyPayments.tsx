import { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import styles from './MyPayments.module.css';
import Button from '@/components/Button/Button';
import api from '@/helpers/API';
import { Payment } from '@/interfaces/payment';
import { PaymentsTable } from '@/pages/MyPayments/PaymetnsTable/PaymetnsTable';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import { PaymentModal } from '@/components/PaymentModal/PaymentModal';
import { formatDate } from '@/helpers/FormatDate';

export function PageMyPayments() {
	const [selectedIds, setSelectedIds] = useState<number[]>([]);
	const [isModalOpen, setIsModalOpen] = useState(false);

	const [payments, setPayments] = useState<Payment[]>([]);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const [limit] = useState(20);
	const [offset, setOffset] = useState(0);
	const [searchParams, setSearchParams] = useSearchParams();

	const [status, setStatus] = useState<'all' | 'paid' | 'unpaid'>(searchParams.get('status') as any || 'all');
	const [searchQuery, setSearchQuery] = useState(searchParams.get('rental_id') || '');
	const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>(searchParams.get('sort') as any || 'asc');

	const selectedPayments = payments.filter(p => selectedIds.includes(p.id));

	const totalAmount = selectedPayments.reduce((sum, p) => sum + p.amount, 0);

	const sortedByStart = [...selectedPayments].sort((a, b) => new Date(a.start_date).getTime() - new Date(b.start_date).getTime());
	const sortedByEnd = [...selectedPayments].sort((a, b) => new Date(b.end_date).getTime() - new Date(a.end_date).getTime());

	const startDate = sortedByStart[0]?.start_date;
	const endDate = sortedByEnd[0]?.end_date;

	useEffect(() => {
		setError("");
		fetchPayments();
	}, [offset, status, searchQuery, sortOrder]);

	useEffect(() => {
		if (error) toast.error(error);
	}, [error]);

	useEffect(() => {
		const params: Record<string, string> = {};
		if (searchQuery.trim()) params.rental_id = searchQuery.trim();
		if (status !== 'all') params.status = status;
		if (sortOrder !== 'desc') params.sort = sortOrder;
		setSearchParams(params);
	}, [searchQuery, status, sortOrder]);

	const fetchPayments = async () => {
		try {
			setIsLoading(true);
			const isPaidParam = status === 'all' ? undefined : status === 'paid';

			const { data } = await api.get<Payment[]>(`/payments/client/`, {
				params: {
					limit,
					offset,
					rental_id: searchQuery.trim() || undefined,
					is_paid: isPaidParam,
					sort_by: 'period_start',
					sort_order: sortOrder,
				}
			});
			setPayments(data);
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.response?.data.error || 'Ошибка загрузки данных');
			}
		} finally {
			setIsLoading(false);
		}
	};

	const handleNext = () => setOffset((prev) => prev + limit);
	const handlePrev = () => setOffset((prev) => Math.max(prev - limit, 0));

	const handleToggleSelect = (id: number) => {
		setSelectedIds((prev) =>
			prev.includes(id) ? prev.filter((i) => i !== id) : [...prev, id]
		);
	};

	const handleToggleSelectAll = (checked: boolean) => {
		if (checked) {
			const unpaidIds = payments.filter(p => !p.status).map(p => p.id);
			setSelectedIds(unpaidIds);
		} else {
			setSelectedIds([]);
		}
	};



	return (
		<>
			<div className={styles.headRow}>
				<h1>Список платёжных операций</h1>
			</div>

			<div className={styles.filters}>
				<input
					type="text"
					placeholder="Поиск по номеру договора..."
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

				{selectedIds.length > 0 && (
					<Button onClick={() => setIsModalOpen(true)}>Оплатить выбранные</Button>
				)}

			</div>

			<div>
				{isLoading ? (
					<div>Загрузка...</div>
				) : (
					<>
						<PaymentsTable
							payments={payments}
							selectedIds={selectedIds}
							onToggleSelect={handleToggleSelect}
							onToggleSelectAll={handleToggleSelectAll}
						/>

						<div className={styles.pagination}>
							<Button className={styles['small']} onClick={handlePrev} disabled={offset === 0}>Назад</Button>
							<span>Показано с {offset + 1} по {offset + payments.length}</span>
							<Button className={styles['small']} onClick={handleNext} disabled={payments.length < limit}>Вперёд</Button>
						</div>
					</>
				)}
			</div>

			<PaymentModal
				isOpen={isModalOpen}	
				paymentIds={selectedIds}
				totalAmount={totalAmount}
				periodStart={formatDate(startDate)}
				periodEnd={formatDate(endDate)}
				onClose={() => setIsModalOpen(false)}
				onSuccess={() => {
					setSelectedIds([]);
					fetchPayments();
				}}
			/>
		</>
	);
}

export default PageMyPayments;
