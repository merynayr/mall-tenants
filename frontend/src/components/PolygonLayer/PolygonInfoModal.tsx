import React, { useState } from 'react';
import styles from './PolygonInfoModal.module.css';

interface PolygonInfoModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (info: { premiseCode: number; note: string }) => void;
}

export const PolygonInfoModal: React.FC<PolygonInfoModalProps> = ({
	isOpen,
	onClose,
	onSave
}) => {
	const [premiseCode, setPremiseCode] = useState<string>('');
	const [note, setNote] = useState<string>('');

	if (!isOpen) return null;

	const handleSave = () => {
		const numericCode = Number(premiseCode);
		if (!premiseCode || isNaN(numericCode) || numericCode <= 0) {
			alert('Пожалуйста, введите корректный код помещения.');
			return;
		}

		onSave({ premiseCode: numericCode, note });
		setPremiseCode('');
		setNote('');
	};

	return (
		<div className={styles.overlay} onClick={onClose}>
			<div className={styles.modal} onClick={e => e.stopPropagation()}>
				<h2>Информация о помещении</h2>

				<div className={styles.inputGroup}>
					<label>Код помещения:</label>
					<input
						type="number"
						value={premiseCode}
						onChange={e => setPremiseCode(e.target.value)}
						placeholder="Введите код помещения"
					/>
				</div>

				<div className={styles.inputGroup}>
					<label>Примечание:</label>
					<textarea
						value={note}
						onChange={e => setNote(e.target.value)}
						placeholder="Введите примечание"
						rows={4}
					/>
				</div>

				<div className={styles.actions}>
					<button className={styles.save} onClick={handleSave}>Сохранить</button>
					<button className={styles.cancel} onClick={onClose}>Отмена</button>
				</div>
			</div>
		</div>
	);
};

export default PolygonInfoModal;
