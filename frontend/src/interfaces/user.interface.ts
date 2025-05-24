import { Role } from '@/types/Role';

export interface Profile {
	user_id: number;
	email: string;
	role: Role;
	address: string;
	name: string;
	phone: string;
}