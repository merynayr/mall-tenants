export type Point = { x: number; y: number };

export type Polygon = {
  floor: number;
  points: Point[];
  label: string;
  premiseCode: number;
};

export type PolygonFromAPI = {
  floor: number;
  points: string;
  label: string;
  premise_code: number;
};
