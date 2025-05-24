import styles from './Premises.module.css';

interface PaymentModalProps {
  amount: number;
  selectedPayment: string | null;
  onSelectPayment: (method: string) => void;
  onCancel: () => void;
  onConfirm: () => void;
}

export function PaymentModal({
  amount,
  selectedPayment,
  onSelectPayment,
  onCancel,
  onConfirm
}: PaymentModalProps) {
  return (
    <div className={styles['modal-overlay']}>
      <div className={styles['modal']}>
        <h2>Выберите способ оплаты</h2>
        <div className={styles['payment-options']}>
          {['card', 'sbp', 'yoomoney'].map((method) => (
            <button
              key={method}
              className={`${styles['payment-button']} ${selectedPayment === method ? styles.selected : ''}`}
              onClick={() => onSelectPayment(method)}
            >
              {method === 'card' && '💳 Банковская карта'}
              {method === 'sbp' && '📱 СБП'}
              {method === 'yoomoney' && '💰 ЮMoney'}
            </button>
          ))}
        </div>

        <p style={{ marginTop: '1rem' }}>
          Сумма к оплате: <strong>{amount} ₽</strong>
        </p>

        <div className={styles['modal-actions']}>
          <button onClick={onCancel}>Отмена</button>
          <button onClick={onConfirm} disabled={!selectedPayment}>Подтвердить</button>
        </div>
      </div>
    </div>
  );
}
