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
						<strong>Обработана:</strong> {a.isProcessed ? 'Да' : 'Нет'}
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
