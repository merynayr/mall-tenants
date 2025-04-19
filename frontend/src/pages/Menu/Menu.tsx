import axios, { AxiosError } from 'axios';
import React, { useState, useEffect } from 'react';
import styles from './Menu.module.css';
import { ControlPanel } from '@/components/PolygonLayer/ControlPanel';
import { DrawingOverlay } from '@/components/PolygonLayer/DrawingOverlay';
import { PolygonInfoModal } from '@/components/PolygonLayer/PolygonInfoModal';
import { PolygonLayer } from '@/components/PolygonLayer/PolygonLayer';
import { PREFIX } from '@/helpers/API';
import { Point, Polygon, PolygonFromAPI } from '@/interfaces/floorplan';

export function PremisesMapperPage() {
	const [isDrawing, setIsDrawing] = useState(false);
	const [currentPoints, setCurrentPoints] = useState<Point[]>([]);
	const [polygons, setPolygons] = useState<Polygon[]>([]);
	const [svgData, setSvgData] = useState<string | null>(null);
	const [floor, setFloor] = useState<number>(1);
	const [showModal, setShowModal] = useState(false);

	useEffect(() => {
		setSvgData('');
		setPolygons([]);
		
		fetchFloorPlan(floor);
		fetchPolygons(floor);
	}, [floor]);

	const fetchFloorPlan = async (floor: number) => {
		try {
			const { data } = await axios.get<string>(`${PREFIX}/floor-plan/${floor}`);
			setSvgData(data);
		} catch (e) {
			if (e instanceof AxiosError) console.error(e.message);
		}
	};

	const fetchPolygons = async (floor: number) => {
		try {
			const { data } = await axios.get<PolygonFromAPI[]>(`${PREFIX}/floor-plan/polygons/${floor}`);
			const parsed = data.map((poly) => ({
				points: poly.points.split(' ').map((pair) => {
					const [x, y] = pair.split(',').map(Number);
					return { x, y };
				}),
				floor: poly.floor,
				label: poly.label,
				premiseCode: poly.premise_code
			}));
			setPolygons(parsed);
		} catch (e) {
			if (e instanceof AxiosError) console.error(e.message);
		}
	};

	const savePolygon = async (polygon: Polygon): Promise<boolean> => {
		try {
			console.log('Текущий этаж при сохранении:', floor);
			await axios.post(`${PREFIX}/floor-plan/polygons/${polygon.premiseCode}`, {
				floor: polygon.floor,
				points: polygon.points.map(p => `${p.x},${p.y}`).join(' '),
				label: polygon.label
			});
			return true;
		} catch (e) {
			if (e instanceof AxiosError) console.error(e.message);
			return false;
		}
	};

	const handleSvgClick = (e: React.MouseEvent<SVGSVGElement>) => {
		if (!isDrawing) return;
		const rect = e.currentTarget.getBoundingClientRect();
		const x = e.clientX - rect.left;
		const y = e.clientY - rect.top;
		setCurrentPoints([...currentPoints, { x, y }]);
	};

	const handleSvgRightClick = (e: React.MouseEvent<SVGSVGElement>) => {
		if (!isDrawing) return;
		e.preventDefault();

		if (currentPoints.length > 0) {
			setCurrentPoints(currentPoints.slice(0, -1));
		}
	};

	const finishDrawing = async () => {
		if (currentPoints.length >= 3) {
			setShowModal(true);
		}
	};

	const handleSavePolygonInfo = async (info: { premiseCode: number; note: string }) => {
		if (info.premiseCode <= 0) {
			alert('Введите корректный код помещения');
			return;
		}
		const newPolygon = {
			floor: floor,
			points: currentPoints,
			label: info.note,
			premiseCode: info.premiseCode
		};
		const isSaved = await savePolygon(newPolygon);
		if (isSaved) {
			setPolygons([...polygons, newPolygon]);
			setCurrentPoints([]);
			setIsDrawing(false);
			setShowModal(false);
		} else {
			alert('Ошибка при сохранении полигона. Попробуйте снова.');
		}
	};

	const resetCurrent = () => {
		setCurrentPoints([]);
		setIsDrawing(false);
	};

	return (
		<div className={styles.wrapper}>
			<h1 className={styles.title}>Редактирование плана этажа</h1>

			<ControlPanel
				isDrawing={isDrawing}
				onStart={() => setIsDrawing(true)}
				onFinish={finishDrawing}
				onReset={resetCurrent}
				floor={floor}
				onFloorChange={(val) => setFloor(val)}
			/>

			<svg
				className={styles.canvas}
				width={800}
				height={500}
				onClick={handleSvgClick}
				onContextMenu={handleSvgRightClick} 
			>
				{svgData && (
					<image
						key={`floor-${floor}`}
						href={`data:image/svg+xml;utf8,${encodeURIComponent(svgData)}`}
						x={0}
						y={0}
						width={800}
						height={500}
						preserveAspectRatio="xMidYMid meet"
					/>
				)}

				<PolygonLayer polygons={polygons} />
				<DrawingOverlay currentPoints={currentPoints} />
			</svg>

			<PolygonInfoModal
				isOpen={showModal}
				onClose={() => setShowModal(false)}
				onSave={handleSavePolygonInfo}
			/>
		</div>
	);
}

export default PremisesMapperPage;