import { Point, Polygon } from '@/interfaces/floorplan';

interface PolygonLayerProps {
  polygons: Polygon[];
	selectedIndex: number | null;
  onSelect: (index: number) => void;
	onDoubleClick: (premiseCode: number) => void;
}

export const PolygonLayer: React.FC<PolygonLayerProps> = ({
	polygons, 
	selectedIndex, 
	onSelect, 
	onDoubleClick 
}) => {
	const pointsString = (pts: Point[]) => pts.map((p) => `${p.x},${p.y}`).join(' ');

	const textLabelPosition = (points: Point[]) => {
		if (points.length === 0) return null;
		const x = points.reduce((sum, p) => sum + p.x, 0) / points.length;
		const y = points.reduce((sum, p) => sum + p.y, 0) / points.length;
		return { x, y };
	};

	const getFillColor = (status: string, isSelected: boolean): string => {
		if (isSelected) return 'rgba(0,123,255,0.3)';
		switch (status) {
		case 'available':
			return 'rgba(40,167,69,0.3)'; 
		case 'occupied':
			return 'rgba(220,53,69,0.3)'; 
		case 'maintenance':
			return 'rgba(255,193,7,0.3)'; 
		default:
			return 'rgba(108,117,125,0.3)';
		}
	};
	
	const getStrokeColor = (status: string, isSelected: boolean): string => {
		if (isSelected) return '#007bff';
		switch (status) {
		case 'available':
			return '#28a745';
		case 'occupied':
			return '#dc3545';
		case 'maintenance':
			return '#ffc107';
		default:
			return '#6c757d';
		}
	};
	
	return (
		<>
			{polygons.map((poly, i) => (
				<g key={i}>
					<polygon
						points={pointsString(poly.points)}
						fill={getFillColor(poly.status, i === selectedIndex)}
						stroke={getStrokeColor(poly.status, i === selectedIndex)}
						strokeWidth={2}
						onClick={(e) => {
							e.stopPropagation();
							onSelect(i);
						}}
						onDoubleClick={() => onDoubleClick(poly.premiseCode)} 
						style={{ cursor: 'pointer' }}
					/>
					{textLabelPosition(poly.points) && (
						<text
							x={textLabelPosition(poly.points)!.x}
							y={textLabelPosition(poly.points)!.y}
							fontSize={14}
							fill="black"
							textAnchor="middle"
						>
							{poly.label}
						</text>
					)}
				</g>
			))}
		</>
	);
};
