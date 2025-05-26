export interface Application {
	id: number;
	organizationName: string;
	contactPerson: string;
	address: string;
	phone: string;
	requisites: string;
	email: string;
	premiseNumber: string;
	additionalInfo?: string;
	isProcessed: boolean;
	createdAt: string;
}
