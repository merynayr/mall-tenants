import { Point } from '@/interfaces/floorplan';

interface DrawingOverlayProps {
  currentPoints: Point[];
}

export const DrawingOverlay: React.FC<DrawingOverlayProps> = ({ currentPoints }) => {
	const pointsString = (pts: Point[]) => pts.map((p) => `${p.x},${p.y}`).join(' ');

	return (
		<>
			{currentPoints.length >= 2 && (
				<polyline
					points={pointsString(currentPoints)}
					fill="none"
					stroke="#00f"
					strokeWidth={2}
					strokeDasharray="4 2"
				/>
			)}
			{currentPoints.map((pt, i) => (
				<circle key={i} cx={pt.x} cy={pt.y} r={4} fill="#00f" />
			))}
		</>
	);
};