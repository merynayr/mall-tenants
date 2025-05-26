import { AxiosError } from 'axios';
import { useEffect, useState } from 'react';
import styles from './Premises.module.css';
import Button from '@/components/Button/Button';
import { PremiseCreateModal } from '@/components/Premises/PremiseCreateModal';
import api from '@/helpers/API';
import { useHasRole } from '@/hooks/Role';
import { Premises } from '@/interfaces/premises';
import { PremiseList } from '@/pages/Premise/PremisesList/PremisesList';
import { toast } from 'react-toastify';


export function PagePremises() {
	const [premises, setPremiseses] = useState<Premises[]>([]);
	const [isLoading, setIsLoading] = useState<boolean>(false);
	const [error, setError] = useState<string | undefined>();
	const [isModalOpen, setModalOpen] = useState(false);


	useEffect(() => {
		getPremises();
	}, []);

	useEffect(() => {
		if (error) {
			toast.error(error);
		}
	}, [error]);
	
	const getPremises = async () => {
		try {
			setIsLoading(true);
			const { data } = await api.get<Premises[]>('/premise', {});
			setPremiseses(data);
			setIsLoading(false);
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.response?.data.error);
			}	
			setIsLoading(false);
			return;
		}
	};

	return <>
		<div className={styles.headRow}>
			<h1>Помещения</h1>
			{useHasRole('moderator') && (
				<Button className={styles.createButton} onClick={() => setModalOpen(true)}>Создать помещение</Button>
			)}
		</div>

		<div>
			{!isLoading && premises.length > 0 && <PremiseList premises={premises} />}
			{isLoading && <>Загружаем помещения...</>}
			{!isLoading && premises.length === 0 && !error &&<>Не найдено помещений по запросу</>}
		</div>

		{isModalOpen && (
			<PremiseCreateModal
				onClose={() => setModalOpen(false)}
				onCreated={() => {
					getPremises();
				}}
			/>
		)}

	</>;
}

export default PagePremises;