import { useEffect, useState } from 'react';
import styles from './Rents.module.css';
import Button from '@/components/Button/Button';
import api from '@/helpers/API';
import { Rent } from '@/interfaces/rent';
import { RentsTable } from '@/pages/Rents/RentsTable/RentsTable';

export function PageRents() {
	const [rents, setRents] = useState<Rent[]>([]);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const [limit] = useState(20);
	const [offset, setOffset] = useState(0);
	
	useEffect(() => {
		fetchRents();
	// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [offset]);

	const fetchRents = async () => {
		try {
			setIsLoading(true);
			const { data } = await api.get<Rent[]>('/rental', {
				params: { limit, offset }
			});
			setRents(data);
		} catch (e) {
			console.error(e);
			setError('Ошибка при загрузке договоров');
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

		<div>
			{error && <div className={styles.error}>{error}</div>}
			{isLoading ? (
				<div>Загрузка...</div>
			) : (
				<>
					<RentsTable rents={rents} />
					<div className={styles.pagination}>
						<Button className={styles['small']} onClick={handlePrev} disabled={offset === 0}>Назад</Button>
						<span>Показано с {offset + 1} по {offset + rents.length}</span>
						<Button className={styles['small']} onClick={handleNext} disabled={rents.length < limit}>Вперёд</Button>
					</div>
				</>
			)}
		</div>
	</>;
}

export default PageRents;
