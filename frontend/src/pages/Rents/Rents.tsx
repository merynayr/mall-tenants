import { useEffect, useState } from 'react';
import styles from './Rents.module.css';
import api from '@/helpers/API';
import { Rent } from '@/interfaces/rent';
import { RentsTable } from '@/pages/Rents/RentsTable/RentsTable';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import Pagination from '@/components/Pagination/Pagination';

export function PageRents() {
	const [rents, setRents] = useState<Rent[]>([]);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const [limit] = useState(11);
	const [offset, setOffset] = useState(0);
	const [searchQuery, setSearchQuery] = useState('');
	const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('desc');

	useEffect(() => {
		fetchRents();
	// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [offset, searchQuery, sortOrder]);

	useEffect(() => {
		if (error) {
			toast.error(error);
		}
	}, [error]);
	
	const fetchRents = async () => {
		try {
			setIsLoading(true);
			const { data } = await api.get<Rent[]>('/rental', {
				params: {
					limit,
					offset,
					client: searchQuery.trim() || undefined,
					sort_by: 'created_at',
					sort_order: sortOrder,
				}
			});
			setRents(data);
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
			<h1>Договора аренды</h1>
		</div>

		<div className={styles.filters}>
			<input
				type="text"
				placeholder="Поиск по имени клиента..."
				value={searchQuery}
				onChange={(e) => setSearchQuery(e.target.value)}
			/>

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
					<RentsTable rents={rents} />

					<Pagination
						offset={offset}
						limit={limit}
						total={rents.length}
						onNext={handleNext}
						onPrev={handlePrev}
					/>
				</>
			)}
		</div>
	</>;
}

export default PageRents;
