import { configureStore } from '@reduxjs/toolkit';
import { saveState } from './storage';
import userSlice, { JWT_PERSISTENT_STATE } from './user.slice';

export const store = configureStore({
	reducer: {
		user: userSlice
	}
});

store.subscribe(() => {
	const { user } = store.getState();

	saveState({
		jwt: user.jwt,
		profile: user.profile 
	 }, JWT_PERSISTENT_STATE);
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispath = typeof store.dispatch;