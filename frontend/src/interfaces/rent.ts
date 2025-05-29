export interface Rent {
	rental_id: number;
	space_code: number;
	client_name: number;
	start_date: string;
	end_date: string;
	paid_months: number;
	created_at: string;
	updated_at?: string;
}

export interface Contract {
	contract_id: number;
	rental_id: number;
	file_path: string;
	is_signed: boolean;
	is_active: boolean;
	signed_at?: string;
}

export interface RentalWithContract {
	rental_id: number;
	space_code: number;
	client_id: number;
	start_date: string;
	end_date: string;
	created_at: string;
	contracts: Contract[];
}
