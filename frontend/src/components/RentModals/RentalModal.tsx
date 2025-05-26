import styles from './Premises.module.css';

interface RentModalProps {
  premiseCode: number;
  startDate: string;
  endDate: string;
  dateError: string | null;
  durationMonths: number | null;
  rentPerMonth: number;
  onStartDateChange: (value: string) => void;
  onEndDateChange: (value: string) => void;
  onCancel: () => void;
  onConfirm: () => void; 
}

export function RentalModal({
  premiseCode,
  startDate,
  endDate,
  dateError,
  durationMonths,
  rentPerMonth,
  onStartDateChange,
  onEndDateChange,
  onCancel,
  onConfirm
}: RentModalProps) {
  return (
    <div className={styles['modal-overlay']}>
      <div className={styles['modal']}>
        <h2>Аренда помещения</h2>
        <p><strong>Помещение:</strong> {premiseCode}</p>

        <label>
          Дата начала:
          <input type="date" value={startDate} onChange={(e) => onStartDateChange(e.target.value)} />
        </label>

        <label>
          Дата окончания:
          <input type="date" value={endDate} onChange={(e) => onEndDateChange(e.target.value)} />
        </label>

        {dateError && <p style={{ color: 'red' }}>{dateError}</p>}

        {durationMonths !== null && (
          <>
            <p><strong>Срок аренды:</strong> {durationMonths} мес.</p>
            <p><strong>Аренда/мес:</strong> {rentPerMonth} ₽</p>
            <p><strong>Итого к оплате:</strong> {durationMonths * rentPerMonth} ₽</p>
          </>
        )}

        <div className={styles['modal-actions']}>
          <button onClick={onCancel}>Отмена</button>
          <button onClick={onConfirm} disabled={!durationMonths}>Арендовать</button>
        </div>
      </div>
    </div>
  );
}
