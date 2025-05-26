import { AxiosError } from 'axios';
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import  './Menu.module.css';
import { FloorPlanCanvas } from '@/components/FloorPlan/FloorPlanCanvas';
import { FloorPlanUploadModal } from '@/components/FloorPlan/FloorPlanUploadModal';
import { ControlPanel } from '@/components/PolygonLayer/ControlPanel';
import { PolygonInfoModal } from '@/components/PolygonLayer/PolygonInfoModal';
import api from '@/helpers/API';
import { Point, Polygon, PolygonFromAPI } from '@/interfaces/floorplan';
import '@/store/storage';
import { toast } from 'react-toastify';

export function PremisesMapperPage() {
	const [isDrawing, setIsDrawing] = useState(false);
	const [currentPoints, setCurrentPoints] = useState<Point[]>([]);
	const [polygons, setPolygons] = useState<Polygon[]>([]);
	const [imageData, setImageData] = useState<string | null>(null);
	const [floor, setFloor] = useState<number>(1);
	const [showModal, setShowModal] = useState(false);
	const [selectedPolygonIndex, setSelectedPolygonIndex] = useState<number | null>(null);
	const [showAddFloorPlanModal, setShowAddFloorPlanModal] = useState(false);
	const [error, setError] = useState<string | null>(null);


	useEffect(() => {
		setImageData('');
		setPolygons([]);
		
		fetchFloorPlan(floor);
		fetchPolygons(floor);
	}, [floor]);

	useEffect(() => {
		if (error) {
			toast.error(error);
		}
	}, [error]);
	
	const navigate = useNavigate();

	const handlePolygonDoubleClick = (premiseCode: number) => {
		navigate(`/premise/${premiseCode}`);
	};

	const handleAddFloorPlan = () => {
		setShowAddFloorPlanModal(true);
	};

	const handleCloseModal = () => {
		setShowAddFloorPlanModal(false);
	};

	const handleUploaded = () => {
		setShowAddFloorPlanModal(false);
	};

	const finishDrawing = async () => {
		if (currentPoints.length >= 3) {
			setShowModal(true);
		}
	};		

	const resetCurrent = () => {
		setCurrentPoints([]);
		setIsDrawing(false);
	};

	const fetchFloorPlan = async (floor: number) => {
		try {
			const response = await api.get(`/floor-plan/${floor}`, {
				responseType: 'blob',
			});

			const blob = response.data as Blob;
			const reader = new FileReader();
			reader.onloadend = () => {
				const base64 = reader.result as string;
				setImageData(base64);
			};
			reader.readAsDataURL(blob);
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.response?.data.error);
			}
		}
	};



	const fetchPolygons = async (floor: number) => {
		try {
			const { data } = await api.get<PolygonFromAPI[]>(`/floor-plan/polygons/${floor}`);
			const parsed = data.map((poly) => ({
				points: poly.points.split(' ').map((pair) => {
					const [x, y] = pair.split(',').map(Number);
					return { x, y };
				}),
				floor: poly.floor,
				label: poly.label,
				premiseCode: poly.premiseCode,
				status: poly.status
			}));
			setPolygons(parsed);
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.response?.data.error);
			}
		}
	};

	const savePolygon = async (polygon: Polygon): Promise<boolean> => {
		try {
			await api.post(`/floor-plan/polygons/${polygon.premiseCode}`, {
				floor: polygon.floor,
				points: polygon.points.map(p => `${p.x},${p.y}`).join(' '),
				label: polygon.label
			});
			return true;
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.response?.data.error);
			}
			return false;
		}
	};

	const DeletePolygon = async () => {
		if (selectedPolygonIndex === null) return;
	
		const polygonToDelete = polygons[selectedPolygonIndex];
	
		const confirmDelete = window.confirm(`Удалить полигон с кодом помещения ${polygonToDelete.premiseCode}?`);
		if (!confirmDelete) return;
	
		try {
			await api.delete(`/floor-plan/polygons/${polygonToDelete.premiseCode}`);
			const newPolygons = [...polygons];
			newPolygons.splice(selectedPolygonIndex, 1);
			setPolygons(newPolygons);
			setSelectedPolygonIndex(null);
		} catch (e) {
			console.error(e);
			if (e instanceof AxiosError) {
				setError(e.response?.data.error);
			}
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

	const handleSavePolygonInfo = async (info: { premiseCode: number; note: string }) => {
		if (info.premiseCode <= 0) {
			alert('Введите корректный код помещения');
			return;
		}
		const newPolygon = {
			floor: floor,
			points: currentPoints,
			label: info.note,
			premiseCode: info.premiseCode,
			status: ''
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


	return (
		<div className="mapper-page">
			<h1>План этажа</h1>

			<div className="mapper-layout">
				<ControlPanel
					isDrawing={isDrawing}
					onStart={() => setIsDrawing(true)}
					onFinish={finishDrawing}
					onReset={resetCurrent}
					floor={floor}
					onFloorChange={(val) => setFloor(val)}
					DeletePolygon={DeletePolygon}
					selectedPolygonIndex={selectedPolygonIndex}
					onAddFloorPlan={handleAddFloorPlan}
				/>

				<FloorPlanCanvas
					imageData={imageData}
					polygons={polygons}
					currentPoints={currentPoints}
					selectedPolygonIndex={selectedPolygonIndex}
					onSvgClick={(e) => {
						handleSvgClick(e);
						setSelectedPolygonIndex(null);
					}}
					onSvgRightClick={handleSvgRightClick}
					onPolygonSelect={setSelectedPolygonIndex}
					onPolygonDoubleClick={handlePolygonDoubleClick}
					floor={floor}
				/>
			</div>

			<PolygonInfoModal
				isOpen={showModal}
				onClose={() => setShowModal(false)}
				onSave={handleSavePolygonInfo}
			/>

			<FloorPlanUploadModal
				isOpen={showAddFloorPlanModal}
				onClose={handleCloseModal}
				onUploaded={handleUploaded}
			/>
		</div>
	);
}

export default PremisesMapperPage;