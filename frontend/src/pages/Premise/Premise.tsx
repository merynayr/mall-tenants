import axios, { AxiosError } from 'axios';
import { useEffect, useState } from 'react';
import styles from './Premises.module.css';
import { PREFIX } from '@/helpers/API';
import { Premises } from '@/interfaces/premises';
import { PremiseList } from '@/pages/Premise/PremisesList/PremisesList';


export function PagePremises() {
	const [premises, setPremiseses] = useState<Premises[]>([]);
	const [isLoading, setIsLoading] = useState<boolean>(false);
	const [error, setError] = useState<string | undefined>();


	useEffect(() => {
		getPremises();
	}, []);


	const getPremises = async () => {
		try {
			setIsLoading(true);
			const { data } = await axios.get<Premises[]>(`${PREFIX}/premise`, {});
			setPremiseses(data);
			setIsLoading(false);
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.message);
			}
			setIsLoading(false);
			return;
		}
	};



	return <>
		<div className={styles['head']}>
			<h1>Помещения</h1>
		</div>
		<div>
			{error && <>{error}</>}
			{!isLoading && premises.length > 0 && <PremiseList premises={premises} />}
			{isLoading && <>Загружаем помещения...</>}
			{!isLoading && premises.length === 0 && <>Не найдено помещений по запросу</>}
		</div>
	</>;
}

export default PagePremises;