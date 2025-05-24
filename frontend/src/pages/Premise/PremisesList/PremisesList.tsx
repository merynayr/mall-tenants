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

	const sortedFloors: number[] = Object.keys(groupedByFloor)
		.map(Number)
		.sort((a, b) => a - b);

	return (
		<div>
			{sortedFloors.map((floorNumber) => {
				const floorPremises = groupedByFloor[floorNumber].sort((a, b) => a.code - b.code);

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
