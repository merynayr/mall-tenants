import React from 'react';
import styles from './PaymetnsTable.module.css';
import { Payment } from '@/interfaces/payment';
import { calculateRentalDuration } from '@/helpers/CountMonth';

interface Props {
  payments: Payment[];
  selectedIds: number[];
  onToggleSelect: (id: number) => void;
  onToggleSelectAll: (checked: boolean) => void;
}

export const PaymentsTable: React.FC<Props> = ({ payments, selectedIds, onToggleSelect, onToggleSelectAll }) => {
  const unpaidPayments = payments.filter(p => !p.status);
  const allSelected = unpaidPayments.every(p => selectedIds.includes(p.id));
  const someSelected = unpaidPayments.some(p => selectedIds.includes(p.id));

  return (
    <table className={styles.table}>
      <thead>
        <tr>
          <th>
            <input
              type="checkbox"
              checked={allSelected}
              onChange={(e) => onToggleSelectAll(e.target.checked)}
              ref={(el) => {
                if (el) el.indeterminate = !allSelected && someSelected;
              }}
              title="Выделить все неоплаченные"
            />
          </th>
          <th>ID</th>
          <th>Договор</th>
          <th>Объект аренды</th>
          <th>Клиент</th>
          <th>Период</th>
          <th>Сумма ₽</th>
          <th>Статус</th>
          <th>Оплата с</th>
          <th>Оплата до</th>
          <th>Дата оплаты</th>
        </tr>
      </thead>
      <tbody>
        {payments.map((payment) => (
          <tr key={payment.id}>
            <td>
              <input
                type="checkbox"
                checked={selectedIds.includes(payment.id)}
                onChange={() => onToggleSelect(payment.id)}
                disabled={payment.status}
              />
            </td>
            <td>{payment.id}</td>
            <td>{payment.rental_id}</td>
            <td>{payment.space_code}</td>
            <td>{payment.client_name}</td>
            <td>{calculateRentalDuration(payment.start_date, payment.end_date)} мес.</td>
            <td>{payment.amount}</td>
            <td>{payment.status ? 'Оплачено' : 'Не оплачено'}</td>
            <td>{new Date(payment.start_date).toLocaleDateString()}</td>
            <td>{new Date(payment.end_date).toLocaleDateString()}</td>
            <td>{payment.payment_date ? new Date(payment.payment_date).toLocaleDateString() : ''}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};
