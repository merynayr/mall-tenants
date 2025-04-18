import { PremisesListProps } from './PremiseListProps';
import styles from './PremisesList.module.css';
import PremiseCard from '@/components/PremisesCard/PremisesCard';
import { Premises } from '@/interfaces/premises';

export function PremiseList({ premises }: PremisesListProps) {
	const groupedByFloor = premises.reduce((acc, premise) => {
		if (!acc[premise.floor]) {
			acc[premise.floor] = [];
		}
		acc[premise.floor].push(premise);
		return acc;
	}, {} as Record<number, Premises[]>);

	return (
		<div>
			{Object.keys(groupedByFloor).map((floor) => {
				const floorNumber = Number(floor);
				const floorPremises = groupedByFloor[floorNumber];

				return (
					<div key={floorNumber}>
						<h3>Этаж {floorNumber}</h3>
						<div className={styles.wrapper}>
							{floorPremises.map((premise) => (
								<PremiseCard
									key={premise.code}
									code={premise.code}
									floor={premise.floor}
									area={premise.area}
									status={premise.status}
								/>
							))}
						</div>
					</div>
				);
			})}
		</div>
	);
}
