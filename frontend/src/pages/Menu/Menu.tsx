import axios, { AxiosError } from 'axios';
import React, { useState, useEffect } from 'react';
import { PREFIX } from '@/helpers/API';

type Point = { x: number; y: number };

type Polygon = {
	points: Point[];
	label: string;
	premiseCode: number;
};

type PolygonFromAPI = {
	points: string;
	label: string;
	premise_code: number;
};

export function DrawTesterPage() {
	const [isDrawing, setIsDrawing] = useState(false);
	const [currentPoints, setCurrentPoints] = useState<Point[]>([]);
	const [polygons, setPolygons] = useState<Polygon[]>([]);
	const [svgData, setSvgData] = useState<string | null>(null);
	const [floor, setFloor] = useState<number>(1);
	const [labelInput, setLabelInput] = useState('');
	const [premiseCodeInput, setPremiseCodeInput] = useState<number>(0);

	const fetchFloorPlan = async (floor: number) => {
		try {
			const { data } = await axios.get<string>(`${PREFIX}/floor-plan/${floor}`);
			setSvgData(data);
		} catch (e) {
			if (e instanceof AxiosError) {
				console.error(e.message);
			}
		}
	};

	const fetchPolygons = async () => {
		try {
			const { data } = await axios.get<PolygonFromAPI[]>(`${PREFIX}/premise/polygons`);
			const parsed: Polygon[] = data.map((poly) => ({
				points: poly.points.split(' ').map((pair) => {
					const [x, y] = pair.split(',').map(Number);
					return { x, y };
				}),
				label: poly.label,
				premiseCode: poly.premise_code
			}));
			setPolygons(parsed);
		} catch (e) {
			if (e instanceof AxiosError) {
				console.error(e.message);
			}
		}
	};

	useEffect(() => {
		fetchFloorPlan(floor);
		fetchPolygons();
	}, [floor]);

	const handleSvgClick = (e: React.MouseEvent<SVGSVGElement>) => {
		if (!isDrawing) return;
		const rect = e.currentTarget.getBoundingClientRect();
		const x = e.clientX - rect.left;
		const y = e.clientY - rect.top;
		setCurrentPoints([...currentPoints, { x, y }]);
	};

	const finishDrawing = async () => {
		if (currentPoints.length >= 3) {
			const newPolygon: Polygon = {
				points: currentPoints,
				label: labelInput,
				premiseCode: premiseCodeInput
			};
			const isSaved = await savePolygon(newPolygon);

			if (isSaved) {
				setPolygons([...polygons, newPolygon]);
			} else {
				console.log('Ошибка сохранения полигона');
			}
		}
		setCurrentPoints([]);
		setLabelInput('');
		setPremiseCodeInput(0);
		setIsDrawing(false);
	};

	const startDrawing = () => {
		setCurrentPoints([]);
		setIsDrawing(true);
	};

	const resetAll = () => {
		setCurrentPoints([]);
		setLabelInput('');
		setPremiseCodeInput(0);
		setIsDrawing(false);
	};
	

	const pointsString = (pts: Point[]) =>
		pts.map((p) => `${p.x},${p.y}`).join(' ');

	const textLabelPosition = (points: Point[]) => {
		if (points.length === 0) return null;
		const x = points.reduce((sum, p) => sum + p.x, 0) / points.length;
		const y = points.reduce((sum, p) => sum + p.y, 0) / points.length;
		return { x, y };
	};

	const savePolygon = async (polygon: Polygon): Promise<boolean> => {
		try {
			await axios.post(`${PREFIX}/premise/${polygon.premiseCode}/polygons`, {
				points: pointsString(polygon.points),
				label: polygon.label
			});
			return true;
		} catch (e) {
			if (e instanceof AxiosError) {
				console.error(e.message);
			}
			return false;
		}
	};

	return (
		<div>
			<h1>Эксперимент: нарисуй области</h1>
			<div style={{ marginBottom: 10 }}>
				<button onClick={startDrawing} disabled={isDrawing}>Начать рисовать</button>
				<button onClick={finishDrawing} disabled={!isDrawing}>Завершить фигуру</button>
				<button onClick={resetAll}>Сбросить текущую фигуру</button>
			</div>

			<div style={{ marginBottom: 10 }}>
				<label>Выберите этаж: </label>
				<input
					type="number"
					value={floor}
					onChange={(e) => setFloor(Number(e.target.value))}
				/>
			</div>
			{isDrawing && (
				<div style={{ marginBottom: 10 }}>
					<label>Подпись (label): </label>
					<input
						type="text"
						value={labelInput}
						onChange={(e) => setLabelInput(e.target.value)}
					/>
					<label> Код помещения: </label>
					<input
						type="number"
						value={premiseCodeInput}
						onChange={(e) => setPremiseCodeInput(Number(e.target.value))}
					/>
				</div>
			)}
			<svg
				width={800}
				height={500}
				style={{
					border: '1px solid #ccc',
					cursor: isDrawing ? 'crosshair' : 'default'
				}}
				onClick={handleSvgClick}
			>
				{svgData && (
					<image
						href={`data:image/svg+xml;utf8,${encodeURIComponent(svgData)}`}
						x={0}
						y={0}
						width={800}
						height={500}
						preserveAspectRatio="xMidYMid meet"
					/>
				)}

				{/* Сохранённые полигоны */}
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

				{/* Текущий незавершённый полигон */}
				{currentPoints.length >= 2 && (
					<polyline
						points={pointsString(currentPoints)}
						fill="none"
						stroke="#00f"
						strokeWidth={2}
						strokeDasharray="4 2"
					/>
				)}

				{/* Точки */}
				{currentPoints.map((pt, i) => (
					<circle key={i} cx={pt.x} cy={pt.y} r={4} fill="#00f" />
				))}
			</svg>
		</div>
	);
}

export default DrawTesterPage;
