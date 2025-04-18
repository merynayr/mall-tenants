import axios from 'axios';
import { useEffect, useState } from 'react';
import { useParams, useLocation } from 'react-router-dom';
import styles from './Premises.module.css';
import { PREFIX } from '@/helpers/API';
import { Premises } from '@/interfaces/premises';

export function PremiseInfo() {
	const { id } = useParams<{ id: string }>();
	const location = useLocation();
	const initialPremise = location.state?.premise as Premises | undefined;

	const [premise, setPremise] = useState<Premises | null>(initialPremise || null);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		if (!premise && id) {
			setIsLoading(true);
			axios.get<Premises>(`${PREFIX}/premise/${id}`)
				.then(({ data }) => setPremise(data))
				.catch((err) => {
					console.error(err);
					setError('Не удалось загрузить данные о помещении');
				})
				.finally(() => setIsLoading(false));
		}
	}, [id, premise]);

	const getReadableType = (type: string) => {
		switch (type) {
		case 'retail':
			return 'Торговое помещение';
		case 'service':
			return 'Техническое помещение';
		default:
			return type;
		}
	};
	
	const getReadableStatus = (status: string | null | undefined) => {
		switch (status) {
		case 'available':
			return 'Свободно';
		case 'occupied':
			return 'Занято';
		case 'maintenance':
			return 'На обслуживании';
		case '':
		case null:
		case undefined:
			return '—';
		default:
			return status;
		}
	};
	
	return (
		<div className={styles['head']}>
			<h1>Помещение</h1>
			<div>
				{error && <p style={{ color: 'red' }}>{error}</p>}
				{isLoading && <p>Загружаем данные...</p>}
				{!isLoading && premise && (
					<div className={styles['premise-info']}>
						<p><span className={styles.label}>Номер:</span> {premise.code}</p>
						<p><span className={styles.label}>Этаж:</span> {premise.floor}</p>
						<p><span className={styles.label}>Площадь:</span> {premise.area} м²</p>
						<p><span className={styles.label}>Тип:</span> {getReadableType(premise.type)}</p>
						<p><span className={styles.label}>Статус:</span> {getReadableStatus(premise.status)}</p>
						<p><span className={styles.label}>Аренда/мес:</span> {premise.rent_per_month || '—'}</p>
						<p><span className={styles.label}>Охрана:</span> {premise.security_system || '—'}</p>
						<p><span className={styles.label}>Кондиционирование:</span> {premise.air_conditioning || '—'}</p>
					</div>

				)}
			</div>
		</div>
	);
}

export default PremiseInfo;
