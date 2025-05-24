import { useState } from 'react';
import styles from './LeaveRequestModal.module.css';
import Button from '@/components/Button/Button';
import api from '@/helpers/API';

interface Props {
	premiseNumber: string;
	onClose: () => void;
}

export default function LeaveRequestModal({ premiseNumber, onClose }: Props) {
	const [organizationName, setOrganizationName] = useState('');
	const [contactPerson, setContactPerson] = useState('');
	const [address, setAddress] = useState('');
	const [phone, setPhone] = useState('');
	const [requisites, setRequisites] = useState('');
	const [email, setEmail] = useState('');
	const [additionalInfo, setadditionalInfo] = useState('');

	const [isSubmitting, setIsSubmitting] = useState(false);

	const handleSubmit = async () => {
		const payload = {
			organizationName,
			contactPerson,
			address,
			phone,
			requisites,
			email,
			premiseNumber,       
			additionalInfo: additionalInfo || undefined,
		};
	
		try {
			setIsSubmitting(true);
			console.log(payload);
			await api.post('/applications/', payload);
			alert('Заявка отправлена. Мы свяжемся с вами.');
			onClose();
		} catch (err) {
			console.error(err);
			alert('Ошибка при отправке заявки. Попробуйте позже.');
		} finally {
			setIsSubmitting(false);
		}
	};
	

	return (
		<div className={styles.overlay}>
			<div className={styles.modal}>
				<h2>Оставить заявку</h2>
				<input type="text" placeholder="Название организации" value={organizationName} onChange={(e) => setOrganizationName(e.target.value)} />
				<input type="text" placeholder="Контактное лицо" value={contactPerson} onChange={(e) => setContactPerson(e.target.value)} />
            	<input type="text" placeholder="Телефон" value={phone} onChange={(e) => setPhone(e.target.value)} />
              	<input type="email" placeholder="Электронная почта" value={email} onChange={(e) => setEmail(e.target.value)} />
                <input type="text" placeholder="Адрес" value={address} onChange={(e) => setAddress(e.target.value)} />
				<input type="text" placeholder="Реквизиты" value={requisites} onChange={(e) => setRequisites(e.target.value)} />
				<input type="text" placeholder="Дополнительная информация" value={additionalInfo} onChange={(e) => setadditionalInfo(e.target.value)} />

				<div className={styles.actions}>
					<Button onClick={onClose}>Отмена</Button>
					<Button onClick={handleSubmit} disabled={isSubmitting}>
						{isSubmitting ? 'Отправка...' : 'Отправить'}
					</Button>
				</div>
			</div>
		</div>
	);
}
