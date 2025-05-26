import { useEffect, useState } from 'react';
import styles from './RentCreateModal.module.css';
import api from '@/helpers/API';
import { toast } from 'react-toastify';

interface Props {
	onClose: () => void;
	onCreated: () => void;
}

export const RentCreateModal: React.FC<Props> = ({ onClose, onCreated }) => {
	const [organizationName, setOrganizationName] = useState('');
	const [contactPerson, setContactPerson] = useState('');
	const [address, setAddress] = useState('');
	const [phone, setPhone] = useState('');
	const [requisites, setRequisites] = useState('');
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const [confirmPassword, setConfirmPassword] = useState('');
	const [role, setRole] = useState('rent');
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		if (error) {
			toast.error(error);
		}
	}, [error]);
	
	const handleSubmit = async () => {
		try {
			await api.post('/rents', {
				organization_name: organizationName,
				contact_person: contactPerson,
				address: address,
				phone: phone,
				requisites: requisites,
				email: email,
				password: password,
				confirm_password: confirmPassword,
				role: role
			});
			onCreated();
			onClose(); 
		} catch (e) {
			console.error(e);
			setError('Ошибка при добавлении клиента');
		}
	};

	return (
		<div className={styles.modalOverlay}>
			<div className={styles.modal}>
				<h2>Создать клиента</h2>

				<input
					type="text"
					placeholder="Название организации"
					value={organizationName}
					onChange={(e) => setOrganizationName(e.target.value)}
				/>
				<input
					type="text"
					placeholder="Контактное лицо"
					value={contactPerson}
					onChange={(e) => setContactPerson(e.target.value)}
				/>
				<input
					type="text"
					placeholder="Адрес"
					value={address}
					onChange={(e) => setAddress(e.target.value)}
				/>
				<input
					type="text"
					placeholder="Телефон"
					value={phone}
					onChange={(e) => setPhone(e.target.value)}
				/>
				<input
					type="text"
					placeholder="Реквизиты"
					value={requisites}
					onChange={(e) => setRequisites(e.target.value)}
				/>
				<input
					type="email"
					placeholder="Электронная почта"
					value={email}
					onChange={(e) => setEmail(e.target.value)}
				/>
				<input
					type="password"
					placeholder="Пароль"
					value={password}
					onChange={(e) => setPassword(e.target.value)}
				/>
				<input
					type="confirm_password"
					placeholder="Повтор пароля"
					value={confirmPassword}
					onChange={(e) => setConfirmPassword(e.target.value)}
				/>
				<select value={role} onChange={(e) => setRole(e.target.value)}>
					<option value="rent">Клиент</option>
					<option value="admin">Администратор</option>
				</select>

				<div className={styles.buttons}>
					<button onClick={handleSubmit}>Создать</button>
					<button onClick={onClose}>Отмена</button>
				</div>
			</div>
		</div>
	);
};
