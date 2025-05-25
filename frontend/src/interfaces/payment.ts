export interface Payment {
  id: number;
	rental_id: number;
	space_code: number;
	client_name: number;
  amount: number;
  status: boolean;
	start_date: string;
	end_date: string;
	paid_months: number;
	created_at: string;
	updated_at?: string;
}
