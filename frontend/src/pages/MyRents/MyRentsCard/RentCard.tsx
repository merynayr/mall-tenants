import { useNavigate } from 'react-router-dom';
import { RentalWithContract } from '@/interfaces/rent';
import cn from 'classnames';
import styles from './RentCard.module.css';
import { formatDate } from '@/helpers/FormatDate';
import api from '@/helpers/API';
import { toast } from 'react-toastify';

interface RentCardProps {
	rent: RentalWithContract;
	onRefresh: () => void;
}

const RentCard: React.FC<RentCardProps> = ({ rent, onRefresh }) => {
	const navigate = useNavigate();

	const handlePayClick = () => {
		navigate(`/my-payments?rental_id=${rent.rental_id}&status=unpaid&sort=asc`);
	};

	const handleDownload = async (contractId: number) => {
		try {
			const response = await api.get(`/contracts/${contractId}`, {
				responseType: 'blob',
			});

			const file = new Blob([response.data], { type: 'application/pdf' });
			const url = window.URL.createObjectURL(file);

			const link = document.createElement('a');
			link.href = url;
			link.download = `contract-${contractId}.pdf`;
			document.body.appendChild(link);
			link.click();
			link.remove();

			window.URL.revokeObjectURL(url);
		} catch (error) {
			console.error('Ошибка при загрузке договора:', error);
		}
	};

	const handleCancel = async () => {
		const unsignedContracts = rent.contracts?.filter(contract => !contract.is_signed) || [];

		if (unsignedContracts.length === 0) {
			toast.info("Нет неподписанных договоров для отмены.");
			return;
		}

		const confirmCancel = window.confirm("Вы уверены, что хотите отменить неподписанные договоры?");
		if (!confirmCancel) return;

		try {
			for (const contract of unsignedContracts) {
				await api.delete(`/contracts/${contract.contract_id}`);
			}
			toast.success("Неподписанные договоры успешно отменены.");
			onRefresh();
		} catch (error: any) {
			toast.error("Ошибка при отмене договора: " + (error.response?.data?.error || error.message));
		}
	};


	async function handleSign(contractId: number, rent: RentalWithContract) {
		const confirmSign = window.confirm("Подписать договор?");
		if (!confirmSign) return;

		try {
			const res = await api.post(`/contracts/${contractId}/sign`, rent);

			if (res.status === 200) {
				toast.success("Договор успешно подписан");
				onRefresh();
			} else {
				toast.error("Ошибка при подписании договора");
			}
		} catch (error: any) {
			toast.error("Ошибка: " + (error.response?.data?.error || error.message));
		}
	}




	return (
		<div className={styles['rent-card']}>
			<div className={styles['rent-title']}>
				<span role="img" aria-label="building">🏢</span>
				Помещение: {rent.space_code}
			</div>

			<p className={styles.period}>
				<span role="img" aria-label="calendar">📅</span>&nbsp;
				<span className={styles.date}>{formatDate(rent.start_date)}</span>
				&nbsp;—&nbsp;
				<span className={styles.date}>{formatDate(rent.end_date)}</span>
			</p>

			<div className={styles['rent-footer']}>
				<span className={styles.badge}>Создано: {formatDate(rent.created_at)}</span>
			</div>

			<div className={styles['contract-section']}>
				<h4>Договоры</h4>
				{rent.contracts?.length === 0 && <p>Нет договоров</p>}
				{rent.contracts?.map(contract => (
					<div key={contract.contract_id} className={styles['contract-item']}>
						<span>ID: {contract.contract_id}</span>
						<button
							className={styles['action-btn']}
							onClick={() => handleDownload(contract.contract_id)}
						>
							📄 Скачать
						</button>
						{!contract.is_signed && (
							<button
								className={cn(styles['action-btn'], styles['action-inline'])}
								onClick={() => handleSign(contract.contract_id, rent)}
							>
								✍️ Подписать
							</button>
						)}
						{contract.is_signed && (
							<span className={styles['status-text']}>✅ Подписано</span>
						)}
					</div>
				))}
			</div>

			<div className={styles['rent-actions']}>
				<button className={styles['action-btn']} onClick={handlePayClick}>
					💳 Оплатить
				</button>
				<button className={styles['action-btn']}>🔁 Продлить</button>
				<button className={styles['action-btn']} onClick={handleCancel}>
					❌ Отменить
				</button>
			</div>
		</div>
	);
};

export default RentCard;
