import React from 'react';
import { DrawingOverlay } from '@/components/PolygonLayer/DrawingOverlay';
import { PolygonLayer } from '@/components/PolygonLayer/PolygonLayer';
import { Point, Polygon } from '@/interfaces/floorplan';

interface FloorPlanCanvasProps {
  svgData: string | null;
  polygons: Polygon[];
  currentPoints: Point[];
  selectedPolygonIndex: number | null;
  onSvgClick: (e: React.MouseEvent<SVGSVGElement>) => void;
  onSvgRightClick: (e: React.MouseEvent<SVGSVGElement>) => void;
  onPolygonSelect: (index: number) => void;
  onPolygonDoubleClick: (premiseCode: number) => void;
  floor: number;
}

export const FloorPlanCanvas: React.FC<FloorPlanCanvasProps> = ({
	svgData,
	polygons,
	currentPoints,
	selectedPolygonIndex,
	onSvgClick,
	onSvgRightClick,
	onPolygonSelect,
	onPolygonDoubleClick,
	floor
}) => {
	return (
		<svg
			width={800}
			height={500}
			onContextMenu={onSvgRightClick}
			onClick={onSvgClick}
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

			<PolygonLayer
				polygons={polygons}
				selectedIndex={selectedPolygonIndex}
				onSelect={onPolygonSelect}
				onDoubleClick={onPolygonDoubleClick}
			/>
      
			<DrawingOverlay currentPoints={currentPoints} />
		</svg>
	);
};
