import { useEffect, useState } from 'react';
import styles from './ClientCreateModal.module.css';
import api from '@/helpers/API';
import { Application } from '@/interfaces/applications';
import { toast } from 'react-toastify';

interface Props {
	onClose: () => void;
	initialData?: Partial<Application>;
}

export const ClientCreateModal: React.FC<Props> = ({ onClose, initialData = {} }) => {
	const [organizationName, setOrganizationName] = useState(initialData.organizationName || '');
	const [contactPerson, setContactPerson] = useState(initialData.contactPerson || '');
	const [address, setAddress] = useState(initialData.address || '');
	const [phone, setPhone] = useState(initialData.phone || '');
	const [requisites, setRequisites] = useState(initialData.requisites || '');
	const [email, setEmail] = useState(initialData.email || '');
	const [password, setPassword] = useState('');
	const [confirmPassword, setConfirmPassword] = useState('');
	const [role, setRole] = useState('client');
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		if (error) {
			toast.error(error);
		}
	}, [error]);
	
	const handleSubmit = async () => {
		try {
			await api.post('/clients', {
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
			setError(null);
			toast.success('Клиент успешно создан');
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
					<option value="client">Клиент</option>
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
