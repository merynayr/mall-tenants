import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { useParams, useLocation } from 'react-router-dom';
import styles from './Premises.module.css';
import Button from '@/components/Button/Button';
import api from '@/helpers/API';
import { Premises } from '@/interfaces/premises';
import { RootState } from '@/store/store';
import { PaymentModal } from '@/components/RentModals/PaymentModal';
import { RentalModal } from '@/components/RentModals/RentalModal';
import { useHasRole } from '@/hooks/Role';
import LeaveRequestModal from '@/components/Applications/LeaveRequestModal';
import { AxiosError } from 'axios';

export function PremiseInfo() {
	const { id } = useParams<{ id: string }>();
	const location = useLocation();
	const initialPremise = location.state?.premise as Premises | undefined;
	const { profile } = useSelector((state: RootState) => state.user);

	const [premise, setPremise] = useState<Premises | null>(initialPremise || null);
	const [isLoading, setIsLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);

	const [isModalOpen, setIsModalOpen] = useState(false);
	const [startDate, setStartDate] = useState<string>('');
	const [endDate, setEndDate] = useState<string>(''); 
	const [durationMonths, setDurationMonths] = useState<number | null>(null);
	const [dateError, setDateError] = useState<string | null>(null);
	const [isPaymentModalOpen, setIsPaymentModalOpen] = useState(false);
	const [selectedPayment, setSelectedPayment] = useState<string | null>(null);
	const [isRequestModalOpen, setIsRequestModalOpen] = useState(false);
	const hasRole = useHasRole('client', 'moderator', 'director');
	
	useEffect(() => {
		if (startDate && endDate) {
			const start = new Date(startDate);
			const end = new Date(endDate);
	
			if (start > end) {
				setDateError('Дата начала позже даты окончания');
				setDurationMonths(null);
				return;
			}

			const today = new Date();
			today.setHours(0, 0, 0, 0);

			if (start <= today) {
				setDateError('Дата начала должна быть позже сегодняшнего дня');
				setDurationMonths(null);
				return;
			}

			if (start.getDate() !== end.getDate()) {
				setDateError('Даты должны быть в один и тот же день месяца (например, с 1 по 1)');
				setDurationMonths(null);
				return;
			}
	
			const monthsDiff = (end.getFullYear() - start.getFullYear()) * 12 + (end.getMonth() - start.getMonth());
	
			if (monthsDiff < 1) {
				setDateError('Минимальный срок аренды — 1 месяц');
				setDurationMonths(null);
				return;
			}
	
			setDateError(null);
			setDurationMonths(monthsDiff);
		}
	}, [startDate, endDate]);

	
	useEffect(() => {
		if (!premise && id) {
			setIsLoading(true);
			api.get<Premises>(`/premise/${id}`)
				.then(({ data }) => setPremise(data))
				.catch((e) => {
					console.error(e);
					if (e instanceof AxiosError) {
						setError(e.response?.data.error);
					}					
				})
			.finally(() => setIsLoading(false));
		}
	}, [id, premise]);

	const resetForm = () => {
		setStartDate('');
		setEndDate('');
		setDurationMonths(null);
		setSelectedPayment(null);
		setDateError(null);
	};
	
	const closeAllModals = () => {
		setIsModalOpen(false);
		setIsPaymentModalOpen(false);
		setStartDate('');
		setEndDate('');
		setDurationMonths(null);
		setDateError(null);
		setSelectedPayment(null);
	};
		
	const getReadableType = (type: string) => {
		switch (type) {
		case 'retail':
			return 'Торговое помещение';
		case 'service':
			return 'Техническое помещение';
		default:
			return type;
		}
	};
	
	const getReadableStatus = (status: string | null | undefined) => {
		switch (status) {
		case 'available':
			return 'Свободно';
		case 'occupied':
			return 'Занято';
		case 'maintenance':
			return 'На обслуживании';
		case '':
		case null:
		case undefined:
			return '—';
		default:
			return status;
		}
	};

	const getReadableExtra = (status: string | null | undefined) => {
		switch (status) {
		case 'NO':
			return 'Отсутствует';
		case 'YES':
			return 'Есть';
		case '':
		case null:
		case undefined:
			return '—';
		default:
			return status;
		}
	};

	const handlePaymentSubmit = async () => {
		if (!premise || !durationMonths || !startDate || !endDate || !selectedPayment) return;
		try {
			await api.post('/rental/', {
				space_code: Number(premise.code),
				client_id: profile?.user_id,
				start_date: new Date(startDate).toISOString(),
				end_date: new Date(endDate).toISOString(),
				paid_months: 2,
				created_at: new Date().toISOString()
			});
	
			alert('Оплата прошла успешно!');
			closeAllModals();
		} catch (error) {
			console.error('Ошибка при отправке аренды:', error);
			alert('Ошибка при оплате. Попробуйте позже.');
		}
	};
	
	
	return (
		<>
			<h1>Помещение {premise?.floor} - {premise?.code}</h1><div>
				{error && <p style={{ color: 'red' }}>{error}</p>}
				{isLoading && <p>Загружаем данные...</p>}
				{!isLoading && premise && (
					<div className={styles['premise-info']}>
						<p><span className={styles.label}>Площадь:</span> {premise.area} м²</p>
						<p><span className={styles.label}>Тип:</span> {getReadableType(premise.type)}</p>
						<p><span className={styles.label}>Статус:</span> {getReadableStatus(premise.status)}</p>
						<p><span className={styles.label}>Аренда/мес:</span> {premise.rent_per_month || '—'}</p>
						<p><span className={styles.label}>Охрана:</span> {getReadableExtra(premise.security_system)}</p>
						<p><span className={styles.label}>Кондиционирование:</span> {getReadableExtra(premise.air_conditioning)}</p>
					</div>
				)}

				<div className={styles['rent-button-container']}>
					<Button
						disabled={premise?.status !== 'available'}
						onClick={() => {
							if (hasRole) {
								setIsModalOpen(true);
							} else {
								setIsRequestModalOpen(true);
							}
						}}
					>
						{hasRole ? "Арендовать" : "Оставить заявку"}
					</Button>
				</div>
			</div>

			{isModalOpen && premise && (
				<RentalModal
					premiseCode={premise.code}
					startDate={startDate}
					endDate={endDate}
					dateError={dateError}
					durationMonths={durationMonths}
					rentPerMonth={premise.rent_per_month ?? 0}
					onStartDateChange={setStartDate}
					onEndDateChange={setEndDate}
					onCancel={() => {
						setIsModalOpen(false);
						resetForm();
					}}
					onNext={() => {
						if (durationMonths) {
							setIsModalOpen(false);
							setIsPaymentModalOpen(true);
						}
					}}
				/>
			)}

			{isPaymentModalOpen && premise && (
				<PaymentModal
					amount={premise.rent_per_month * 2}
					selectedPayment={selectedPayment}
					onSelectPayment={setSelectedPayment}
					onCancel={() => {
						setIsPaymentModalOpen(false);
						setIsModalOpen(true);
					}}
					onConfirm={() => {
						if (selectedPayment) {
							handlePaymentSubmit();
							closeAllModals();
						}
					}}
				/>
			)}

			{isRequestModalOpen && premise && (
				<LeaveRequestModal
					premiseNumber={String(premise.code)}
					onClose={() => setIsRequestModalOpen(false)}
				/>
			)}

		</>
	);
}

export default PremiseInfo;
