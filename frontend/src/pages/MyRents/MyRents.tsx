import React, { useEffect, useState } from 'react';
import { RentalWithContract } from '@/interfaces/rent';
import api from '@/helpers/API';
import { RootState } from '@/store/store';
import { useSelector } from 'react-redux';
import RentCard from './MyRentsCard/RentCard';
import styles from './MyRents.module.css'; 
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';

const PageMyRents: React.FC = () => {
	const [rents, setRents] = useState<RentalWithContract[]>([]);
	const [loading, setLoading] = useState(true);
	const { profile } = useSelector((state: RootState) => state.user);
	const [error, setError] = useState<string | null>(null);


	useEffect(() => {
		if (error) {
			toast.error(error);
		}
	}, [error]);
	
	useEffect(() => {
		const fetchRents = async (id: number | undefined) => {
			try {
				const response = await api.get<RentalWithContract[]>(`/rental/${id}`);
				setRents(response.data);
				console.log(response.data)
			} catch (e) {
				console.error(e);
				if (e instanceof AxiosError) {
					setError(e.response?.data.error);
				}
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
