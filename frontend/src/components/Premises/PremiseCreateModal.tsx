import React, { useState } from 'react';
import styles from './PremiseCreateModal.module.css';
import api from '@/helpers/API';

interface Props {
	onClose: () => void;
	onCreated: () => void;
}

export const PremiseCreateModal: React.FC<Props> = ({ onClose, onCreated }) => {
	const [floor, setFloor] = useState('');
	const [premiseCode, setPremiseCode] = useState('');
	const [area, setArea] = useState('');
	const [type, setType] = useState('retail');
	const [status, setStatus] = useState('available');
	const [rent, setRent] = useState('');
	const [security, setSecurity] = useState('YES');
	const [air, setAir] = useState('YES');
	const [error, setError] = useState<string | null>(null);

	const handleSubmit = async () => {
		const codeNum = Number(premiseCode);
		const floorNum = Number(floor);
		const areaNum = Number(area);
		const rentNum = Number(rent);

		if (
			!codeNum || codeNum <= 0 ||
			!floorNum || floorNum < 0 ||
			!areaNum || areaNum <= 0 ||
			!rentNum || rentNum < 0
		) {
			setError('Пожалуйста, заполните все числовые поля корректными значениями');
			return;
		}

		try {
			await api.post('/premise', {
				code: codeNum,
				floor: floorNum,
				area: areaNum,
				type,
				status,
				rent_per_month: rentNum,
				security_system: security,
				air_conditioning: air
			});
			onCreated();
			onClose();
		} catch (e) {
			console.error(e);
			setError('Ошибка при создании помещения');
		}
	};

	return (
		<div className={styles.modalOverlay}>
			<div className={styles.modal}>
				<h2>Создать помещение</h2>

				<input type="number" placeholder="Этаж" value={floor} onChange={(e) => setFloor(e.target.value)} />
				<input type="number" placeholder="Номер" value={premiseCode} onChange={(e) => setPremiseCode(e.target.value)} />

				<div className={styles.inputWithSuffix}>
					<input
						type="number"
						placeholder="Площадь"
						value={area}
						onChange={(e) => setArea(e.target.value)}
					/>
					<span className={styles.suffix}>м²</span>
				</div>

				<select value={type} onChange={(e) => setType(e.target.value)}>
					<option value="retail">Торговое</option>
					<option value="service">Служебное</option>
				</select>

				<select value={status} onChange={(e) => setStatus(e.target.value)}>
					<option value="available">Доступно</option>
					<option value="occupied">Занято</option>
					<option value="maintenance">Обслуживание</option>
				</select>

				<div className={styles.inputWithSuffix}>
					<input
						type="number"
						placeholder="Аренда в месяц"
						value={rent}
						onChange={(e) => setRent(e.target.value)}
					/>
					<span className={styles.suffix}>₽</span>
				</div>

				<select value={security} onChange={(e) => setSecurity(e.target.value)}>
					<option value="YES">С охраной</option>
					<option value="NO">Без охраны</option>
				</select>

				<select value={air} onChange={(e) => setAir(e.target.value)}>
					<option value="YES">Кондиционирование</option>
					<option value="NO">Нет кондиционирования</option>
				</select>

				{error && <p className={styles.error}>{error}</p>}

				<div className={styles.buttons}>
					<button onClick={handleSubmit}>Создать</button>
					<button onClick={onClose}>Отмена</button>
				</div>
			</div>
		</div>
	);
};
