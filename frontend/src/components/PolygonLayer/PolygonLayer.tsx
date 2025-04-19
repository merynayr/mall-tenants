import { Point, Polygon } from '@/interfaces/floorplan';

interface PolygonLayerProps {
  polygons: Polygon[];
}

export const PolygonLayer: React.FC<PolygonLayerProps> = ({ polygons }) => {
	const pointsString = (pts: Point[]) => pts.map((p) => `${p.x},${p.y}`).join(' ');

	const textLabelPosition = (points: Point[]) => {
		if (points.length === 0) return null;
		const x = points.reduce((sum, p) => sum + p.x, 0) / points.length;
		const y = points.reduce((sum, p) => sum + p.y, 0) / points.length;
		return { x, y };
	};

	return (
		<>
			{polygons.map((poly, i) => (
				<g key={i}>
					<polygon
						points={pointsString(poly.points)}
						fill="rgba(0,200,0,0.3)"
						stroke="#080"
						strokeWidth={2}
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
