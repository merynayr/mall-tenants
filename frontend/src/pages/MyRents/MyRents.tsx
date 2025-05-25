import React, { useEffect, useState } from 'react';
import { Rent } from '@/interfaces/rent';
import api from '@/helpers/API';
import { RootState } from '@/store/store';
import { useSelector } from 'react-redux';
import RentCard from './MyRentsCard/RentCard';
import styles from './MyRents.module.css'; 

const PageMyRents: React.FC = () => {
	const [rents, setRents] = useState<Rent[]>([]);
	const [loading, setLoading] = useState(true);
	const { profile } = useSelector((state: RootState) => state.user);

	useEffect(() => {
		const fetchRents = async (id: number | undefined) => {
			try {
				const response = await api.get<Rent[]>(`/rental/${id}`);
				setRents(response.data);
			} catch (error) {
				console.error('Ошибка при загрузке аренд:', error);
			} finally {
				setLoading(false);
			}
		};

		if (profile?.user_id) {
			fetchRents(profile.user_id);
		}
	}, [profile?.user_id]);

	if (loading) return <p>Загрузка...</p>;
	if (!rents.length) return <p>У вас нет аренд.</p>;

	return (
		<>
			<div className={styles.headRow}>
				<h1>Мои аренды</h1>
			</div>
			
			<div className={styles['rents-grid']}>
				{rents.map((rent) => (
					<RentCard key={rent.rental_id} rent={rent} />
				))}
			</div></>
	);
};

export default PageMyRents;
