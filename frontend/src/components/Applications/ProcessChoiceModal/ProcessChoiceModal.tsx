import styles from './ProcessChoiceModal.module.css';

interface Props {
	applicationId: number;
	onClose: () => void;
	onCreateClient: () => void;
	onReject: () => void;
}

export const ProcessChoiceModal: React.FC<Props> = ({ onClose, onCreateClient, onReject }) => {
	return (
		<div className={styles.overlay}>
			<div className={styles.modal}>
				<h2>Что сделать с заявкой?</h2>
				<div className={styles.buttons}>
					<button className={styles.primary} onClick={onCreateClient}>Создать клиента</button>
					<button className={styles.danger} onClick={onReject}>Отклонить</button>
					<button className={styles.secondary} onClick={onClose}>Отмена</button>
				</div>
			</div>
		</div>
	);
};
