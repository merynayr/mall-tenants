import { Application } from '@/interfaces/applications';
import styles from './ApplicationsTable.module.css';

interface Props {
	applications: Application[];
	onProcess: (application: Application) => void;
}

export function ApplicationsTable({ applications, onProcess }: Props) {
	if (applications.length === 0) {
		return <p>Нет заявок для отображения.</p>;
	}	
	console.log(applications)

	return (
		<div className={styles.cardList}>
			{applications.map((a) => (
				<div key={a.id} className={styles.card}>
					<div className={styles.row}>
						<strong>ID:</strong> {a.id}
					</div>
					<div className={styles.row}>
						<strong>Организация:</strong> {a.organizationName}
					</div>
					<div className={styles.row}>
						<strong>Контактное лицо:</strong> {a.contactPerson}
					</div>
					<div className={styles.row}>
						<strong>Телефон:</strong> {a.phone}
					</div>
					<div className={styles.row}>
						<strong>Email:</strong> {a.email}
					</div>
					<div className={styles.row}>
						<strong>№ помещения:</strong> {a.premiseNumber}
					</div>
					<div className={styles.row}>
						<strong>Даты аренды:</strong>{' '}	
						{`с ${new Date(a.startDate).toLocaleDateString('ru-RU')} по ${new Date(a.endDate).toLocaleDateString('ru-RU')}`}
					</div>
					<div className={styles.row}>
						<strong>Обработана:</strong>{' '}
						<span className={a.isProcessed ? styles.processed : styles.unprocessed}>
							{a.isProcessed ? 'Да' : 'Нет'}
						</span>
					</div>
					<div className={styles.row}>
						<strong>Создана:</strong> {new Date(a.createdAt).toLocaleString()}
					</div>
					<div className={styles.row}>
						<strong>Доп. информация:</strong> {a.additionalInfo || '—'}
					</div>
					<div className={styles.actions}>
						<button onClick={() => onProcess(a)}>Обработать</button>
					</div>
				</div>
			))}
		</div>
	);
}
