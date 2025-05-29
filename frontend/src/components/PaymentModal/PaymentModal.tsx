import React, { useState } from 'react';
import Modal from 'react-modal';
import styles from './PaymentModal.module.css';
import Button from '@/components/Button/Button';
import api from '@/helpers/API';
import { toast } from 'react-toastify';

interface PaymentModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSuccess: () => void;
  paymentIds: number[];
  totalAmount: number;
  periodStart: string;
  periodEnd: string;
}

export const PaymentModal: React.FC<PaymentModalProps> = ({
  isOpen,
  onClose,
  onSuccess,
  paymentIds,
  totalAmount,
  periodStart,
  periodEnd
}) => {
  const [method, setMethod] = useState('card');
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async () => {
    setIsSubmitting(true);
    try {
      await api.patch('/payments/mark-paid', {
        payment_ids: paymentIds,
      });
      toast.success('Оплата прошла успешно');
      onSuccess();
      onClose();
    } catch (error) {
      toast.error('Ошибка при оплате');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Modal
      isOpen={isOpen}
      onRequestClose={onClose}
      className={styles.modalContent}
      overlayClassName={styles.modalOverlay}
      contentLabel="Подтверждение оплаты"
      closeTimeoutMS={200}
    >

      <h2 className={styles.title}>Подтверждение оплаты</h2>
      <div className={styles.infoBlock}>
        <div className={styles.infoCard}>
          <div className={styles.infoTitle}>Период оплаты</div>
          <div className={styles.infoValue}>с {periodStart} по {periodEnd}</div>
        </div>
        <div className={styles.infoCard}>
          <div className={styles.infoTitle}>Количество месяцев</div>
          <div className={styles.infoValue}>{paymentIds.length}</div>
        </div>
        <div className={styles.infoCard}>
          <div className={styles.infoTitle}>Сумма к оплате</div>
          <div className={styles.infoValue}>{totalAmount} ₽</div>
        </div>
      </div>

      <div className={styles.field}>
        <label htmlFor="method">Способ оплаты:</label>
        <select id="method" value={method} onChange={(e) => setMethod(e.target.value)}>
          <option value="card">Карта</option>
          <option value="bank">СБП</option>
        </select>
      </div>

      <div className={styles.actions}>
        <Button onClick={handleSubmit} disabled={isSubmitting}>Оплатить</Button>
        <Button onClick={onClose}>Отмена</Button>
      </div>
    </Modal>
  );
};
