import React, { useState } from 'react';
import styles from './FloorPlanUploadModal.module.css';
import api from '@/helpers/API';

interface Props {
	isOpen: boolean;
	onClose: () => void;
	onUploaded: () => void;
}

export const FloorPlanUploadModal: React.FC<Props> = ({ isOpen, onClose, onUploaded }) => {
	const [floor, setFloor] = useState<string>('');
	const [file, setFile] = useState<File | null>(null);
	const [error, setError] = useState<string | null>(null);

	const handleSubmit = async () => {
		const numericCode = Number(floor);
		if (!floor || isNaN(numericCode) || numericCode <= 0) {
			setError('Пожалуйста, введите корректный номер этажа.');
			return;
		}
		if (!file) {
			setError('Выберите файл формата PNG');
			return;
		}
		const validTypes = ['image/png'];
		if (!validTypes.includes(file.type)) {
			setError('Допустимы только файлы формата PNG');
			return;
		}

		const formData = new FormData();
		formData.append('floor', floor);
		formData.append('file', file);

		try {
			await api.post(`/floor-plan/${floor}`, formData, {
				headers: {
					'Content-Type': 'multipart/form-data'
				}
			});
			clearForm();
			onUploaded();
			onClose();
		} catch (err) {
			console.error(err);
			setError('Ошибка при загрузке');
		}
	};

	const handleCancel = () => {
		clearForm();
		onClose();
	};
  
	const clearForm = () => {
		setFloor('');
		setFile(null);
		setError(null);
	};

	if (!isOpen) return null;

	return (
		<div className={styles.modalOverlay}>
			<div className={styles.modal}>
				<h2>Добавить план этажа</h2>

				<label>Номер этажа:</label>
				<input
					type="number"
					value={floor}
					onChange={(e) => setFloor(e.target.value)}
				/>

				<label>SVG файл:</label>
				<input
					type="file"
					accept=".png,image/png"
					onChange={(e) => {
						const f = e.target.files?.[0] || null;
						setFile(f);
					}}
				/>

				{error && <p className={styles.error}>{error}</p>}

				<div className={styles.buttons}>
					<button onClick={handleSubmit}>Загрузить</button>
					<button onClick={handleCancel}>Отмена</button>
				</div>
			</div>
		</div>
	);
};
