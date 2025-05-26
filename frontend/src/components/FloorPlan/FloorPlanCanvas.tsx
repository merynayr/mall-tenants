import React from 'react';
import { DrawingOverlay } from '@/components/PolygonLayer/DrawingOverlay';
import { PolygonLayer } from '@/components/PolygonLayer/PolygonLayer';
import { Point, Polygon } from '@/interfaces/floorplan';
import './FloorPlanCanvas.css';
interface FloorPlanCanvasProps {
  imageData: string | null;
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
  imageData,
  polygons,
  currentPoints,
  selectedPolygonIndex,
  onSvgClick,
  onSvgRightClick,
  onPolygonSelect,
  onPolygonDoubleClick,
  floor
}) => {
  const getPngHref = () => {
    if (!imageData) return null;
    return imageData.startsWith('data:image/') ? imageData : `data:image/png;base64,${imageData}`;
  };

  const href = getPngHref();

  return (
    <div className="floor-plan-container">
      <svg
        width={738}
        height={393}
        onContextMenu={onSvgRightClick}
        onClick={onSvgClick}
        className="floor-plan-svg size"
      >
        {href && (
          <image
            key={`floor-${floor}`}
            href={href}
            x={0}
            y={0}
            className='size'
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
    </div>
  );
};