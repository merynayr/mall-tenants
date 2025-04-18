import React from 'react';
import { Premises } from '@/interfaces/premises';
import './Premises.css';

interface Props {
  premise: Premises;
  onClose: () => void;
}

const PremiseModal: React.FC<Props> = ({ premise, onClose }) => {
	return (
		<div className="modal-overlay" onClick={onClose}>
			<div className="modal-content" onClick={(e) => e.stopPropagation()}>
				<button className="modal-close" onClick={onClose}>×</button>
				<h2>Помещение {premise.floor}{premise.code}</h2>
				<p><strong>Площадь:</strong> {premise.area} м²</p>
				<p><strong>Тип:</strong> {premise.type}</p>
				<p><strong>Статус:</strong> {premise.status}</p>
				<p><strong>Арендная стоимость:</strong> {premise.rent_per_month || '—'}</p>
				<p><strong>Охранная система:</strong> {premise.security_system || '—'}</p>
				<p><strong>Кондиционер:</strong> {premise.air_conditioning || '—'}</p>
			</div>
		</div>
	);
};

export default PremiseModal;
