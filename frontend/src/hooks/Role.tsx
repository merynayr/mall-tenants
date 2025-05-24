import { useSelector } from 'react-redux';
import { RootState } from '@/store/store';

export function useHasRole(...roles: string[]): boolean {
	const { profile } = useSelector((state: RootState) => state.user);
	const jwt = useSelector((s: RootState) => s.user.jwt);
	
	const role = profile?.role?.toString();

	if (!jwt || !role) return false;

	return roles.includes(role);
}
