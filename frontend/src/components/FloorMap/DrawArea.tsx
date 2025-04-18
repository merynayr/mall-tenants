import { useState } from 'react';

export function DrawArea() {
	const [points, setPoints] = useState<{ x: number; y: number }[]>([]);

	const handleClick = (e: React.MouseEvent<SVGSVGElement>) => {
		const rect = e.currentTarget.getBoundingClientRect();
		const x = e.clientX - rect.left;
		const y = e.clientY - rect.top;
		setPoints([...points, { x, y }]);
	};

	const handleReset = () => {
		setPoints([]);
	};

	const pointsString = points.map(p => `${p.x},${p.y}`).join(' ');

	return (
		<div>
			<svg
				width={600}
				height={400}
				style={{ border: '1px solid #ccc', cursor: 'crosshair' }}
				onClick={handleClick}
			>
				{/* нарисованная область */}
				{points.length > 2 && (
					<polygon
						points={pointsString}
						fill="rgba(100, 200, 255, 0.5)"
						stroke="blue"
						strokeWidth={2}
					/>
				)}

				{/* точки */}
				{points.map((point, index) => (
					<circle key={index} cx={point.x} cy={point.y} r={4} fill="red" />
				))}
			</svg>

			<div style={{ marginTop: '1rem' }}>
				<button onClick={handleReset}>Очистить</button>
				<p>
					<code>{pointsString}</code>
				</p>
			</div>
		</div>
	);
}
