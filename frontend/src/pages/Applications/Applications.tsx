import { useEffect, useState } from 'react';
import styles from './Applications.module.css';
import api from '@/helpers/API';
import { Application } from '@/interfaces/applications';
import { AxiosError } from 'axios';
import { ApplicationsTable } from './ApplicationsTable/ApplicationsTable';
import { ClientCreateModal } from '@/components/Clients/ClientCreateModal';
import { ProcessChoiceModal } from '@/components/Applications/ProcessChoiceModal/ProcessChoiceModal';
import Button from '@/components/Button/Button';
import { toast } from 'react-toastify';

export function PageApplications() {
	const [applications, setApplications] = useState<Application[]>([]);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const [offset, setOffset] = useState(0);
	const [limit] = useState(8);
	const [filter, setFilter] = useState<'all' | 'true' | 'false'>('all');

	const [selectedApp, setSelectedApp] = useState<Application | null>(null);
	const [showCreateClientModal, setShowCreateClientModal] = useState(false);
	const [showProcessChoiceModal, setShowProcessChoiceModal] = useState(false);

	useEffect(() => {
		fetchApplications();
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [offset, filter]);

	useEffect(() => {
		if (error) {
			toast.error(error);
		}
	}, [error]);
	
	const fetchApplications = async () => {
		try {
			setIsLoading(true);
			setError(null);
			const params: any = { limit, offset };
			if (filter !== 'all') params.processed = filter;

			const { data } = await api.get<Application[]>('/applications', { params });
			setApplications(data);
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.response?.data.error || 'Ошибка при загрузке');
			} else {
				setError('Произошла неизвестная ошибка');
			}
			setApplications([]);
		} finally {
			setIsLoading(false);
		}
	};


	const handleProcessClick = (app: Application) => {
		setSelectedApp(app);
		setShowProcessChoiceModal(true);
	};

	const handleReject = async () => {
		if (!selectedApp) return;
		try {
			await api.patch(`/applications/${selectedApp.id}/status`, { processed: true });
			fetchApplications();
		} catch (e) {
			console.error(e);
		}
		setShowProcessChoiceModal(false);
	};

	const handleCreateClient = () => {
		setShowProcessChoiceModal(false);
		setShowCreateClientModal(true);
	};

	const handleClientCreated = async () => {
		if (!selectedApp) return;

		try {
			await api.patch(`/applications/${selectedApp.id}/status`, { processed: true });
			await fetchApplications();
		} catch (e) {
			console.error('Ошибка при обновлении статуса заявки', e);
		} finally {
			setShowCreateClientModal(false);
		}
	};

	const handleNext = () => setOffset((prev) => prev + limit);
	const handlePrev = () => setOffset((prev) => Math.max(prev - limit, 0));

	return (
		<div className={styles.page}>
			<h1>Заявки</h1>

			<div className={styles.controls}>
				<label>Фильтр:</label>
				<select value={filter} onChange={(e) => setFilter(e.target.value as any)}>
					<option value="all">Все</option>
					<option value="true">Обработанные</option>
					<option value="false">Необработанные</option>
				</select>
			</div>

		
			{isLoading ? (
				<div>Загрузка...</div>
			) : (
				<>
					<div className={styles.tableWrapper}>
						<ApplicationsTable applications={applications} onProcess={handleProcessClick} />
					</div>
					<div className={styles.pagination}>
						<Button className={styles['small']} onClick={handlePrev} disabled={offset === 0}>
							Назад
						</Button>
						<span>
							Показано с {offset + 1} по {offset + applications.length}
						</span>
						<Button className={styles['small']} onClick={handleNext} disabled={applications.length < limit}>
							Вперёд
						</Button>
					</div>
				</>
			)}


			{showProcessChoiceModal && selectedApp && (
				<ProcessChoiceModal
					applicationId={selectedApp.id}
					onClose={() => setShowProcessChoiceModal(false)}
					onCreateClient={handleCreateClient}
					onReject={handleReject}
				/>
			)}

			{showCreateClientModal && selectedApp && (
				<ClientCreateModal
					initialData={selectedApp}
					onClose={() => setShowCreateClientModal(false)}
					onCreated={handleClientCreated}
				/>
			)}
		</div>
	);
}
