import React from 'react';
import './FloorMap.module.css'; // Стили отдельно

export type PlanPremiseStatus = 'available' | 'occupied' | 'maintenance';

export interface PlanPremise {
  code: number;
  points: string; // координаты для полигона
  status: PlanPremiseStatus;
}


type FloorMapProps = {
  premises: PlanPremise[];
  onSelect: (code: number) => void;
};

export const FloorMap: React.FC<FloorMapProps> = ({ premises, onSelect }) => {
	return (
		<svg viewBox="0 0 800 600" className="floor-map">
			{premises.map((p) => (
				<polygon
					key={p.code}
					points={p.points}
					className={`premise ${p.status}`}
					onClick={() => onSelect(p.code)}
				/>
			))}
		</svg>
	);
};
