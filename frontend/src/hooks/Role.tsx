import { useSelector } from 'react-redux';
import { RootState } from '@/store/store';

export function useHasRole(...roles: string[]): boolean {
	const { profile } = useSelector((state: RootState) => state.user);

	const role = profile?.role?.toString();

	if (!role) return false;

	return roles.includes(role);
}
