export type PremisesType = 'retail' | 'service';

export type PremisesStatus = 'available' | 'occupied' | 'maintenance';

export interface Premises {
  code: number;
  floor: number;
  area: number;
  type: PremisesType;
  rent_per_month: number;
  status: PremisesStatus;
  security_system: string;
  air_conditioning: string;
}
