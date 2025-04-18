import { Link } from 'react-router-dom';
import styles from './PremiseCard.module.css';
import { PremisesProps } from './PremisesCardProps';

function PremiseCard(props: PremisesProps) {

	const getStatusClass = () => {
		switch (props.status) {
		case 'available':
			return styles['premise-free'];
		case 'maintenance':
			return styles['premise-medium'];
		case 'occupied':
			return styles['premise-high'];
		default:
			return '';
		}
	};

	return (
		<Link to={`/premise/${props.code}`} className={styles['link']}>
			<div className={`${styles['premise-item']} ${getStatusClass()}`}>
				<div className="code">
					<strong>{props.floor} - {props.code}</strong>
				</div>
				<div className="area">
					{props.area} м²
				</div>
			</div>
		</Link>
	);
}

export default PremiseCard;
