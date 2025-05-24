import { useEffect, useState } from 'react';
import styles from './Clients.module.css';
import Button from '@/components/Button/Button';
import { ClientCreateModal } from '@/components/Clients/ClientCreateModal';
import api from '@/helpers/API';
import { Client } from '@/interfaces/client';
import { ClientsTable } from '@/pages/Clients/ClientsTable/ClientsTable';
import { useHasRole } from '@/hooks/Role';

export function PageClients() {
	const [clients, setClients] = useState<Client[]>([]);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const [limit] = useState(20);
	const [offset, setOffset] = useState(0);
	const [showModal, setShowModal] = useState(false);
	
	useEffect(() => {
		fetchClients();
	// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [offset]);

	const fetchClients = async () => {
		try {
			setIsLoading(true);
			const { data } = await api.get<Client[]>('/clients', {
				params: { limit, offset }
			});
			setClients(data);
		} catch (e) {
			console.error(e);
			setError('Ошибка при загрузке клиентов');
		} finally {
			setIsLoading(false);
		}
	};

	const handleAddClient = () => {
		setShowModal(true);
	};

	const handleCloseModal = () => {
		setShowModal(false);
	};


	const handleNext = () => setOffset((prev) => prev + limit);
	const handlePrev = () => setOffset((prev) => Math.max(prev - limit, 0));

	return <>
	
		<div className={styles.headRow}>
			<h1>Клиенты</h1>
			{useHasRole('moderator') && (
			<Button className={styles.addButton} onClick={handleAddClient}>Создать клиента</Button>
			)}
		</div>

		<div>
			{error && <div className={styles.error}>{error}</div>}
			{isLoading ? (
				<div>Загрузка...</div>
			) : (
				<>
					<ClientsTable clients={clients} />
					<div className={styles.pagination}>
						<Button className={styles['small']} onClick={handlePrev} disabled={offset === 0}>Назад</Button>
						<span>Показано с {offset + 1} по {offset + clients.length}</span>
						<Button className={styles['small']} onClick={handleNext} disabled={clients.length < limit}>Вперёд</Button>
					</div>

					{showModal && <ClientCreateModal onClose={handleCloseModal} onCreated={fetchClients} />}
				</>
			)}
		</div>
			
	</>;
}

export default PageClients;
