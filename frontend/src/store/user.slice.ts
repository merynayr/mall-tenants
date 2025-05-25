import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { AxiosError } from 'axios';
import { loadState } from './storage';
import { RootState } from './store';
import api from '@/helpers/API';
import { LoginResponse } from '@/interfaces/auth.interface';
import { Profile } from '@/interfaces/user.interface';

export const JWT_PERSISTENT_STATE = 'userData';

export interface UserPersistentState {
	jwt: string | null;
	profile: Profile | null;
}

export interface UserState {
	jwt: string | null;
	loginErrorMessage?: string;
	registerErrorMessage?: string;
	profile?: Profile | null;
}

const initialState: UserState = {
	jwt: loadState<UserPersistentState>(JWT_PERSISTENT_STATE)?.jwt ?? null,
	profile: loadState<UserPersistentState>(JWT_PERSISTENT_STATE)?.profile ?? null
};

// login action
export const login = createAsyncThunk<LoginResponse, { email: string, password: string }>(
	'user/login',
	async (params, { rejectWithValue }) => {
		try {
			const { data } = await api.post<LoginResponse>('/auth/login', {
				email: params.email,
				password: params.password
			});
			return data;
		} catch (e) {
			if (e instanceof AxiosError) {
				return rejectWithValue(e.response?.data.error);
			}
			return rejectWithValue('An unexpected error occurred.');
		}
	}
);

// getProfile action
export const getProfile = createAsyncThunk<Profile, void, { state: RootState }>(
	'user/getProfile',
	async () => {
		const { data } = await api.get<Profile>('/clients/profile');
		return data;
	}
);

const userSlice = createSlice({
	name: 'user',
	initialState,
	reducers: {
		logout: (state) => {
			state.jwt = null;
			state.profile = undefined;
		},
		clearLoginError: (state) => {
			state.loginErrorMessage = undefined;
		},
		clearRegisterError: (state) => {
			state.registerErrorMessage = undefined;
		}
	},
	extraReducers: (builder) => {
		builder.addCase(login.fulfilled, (state, action) => {
			if (!action.payload) return;
			state.jwt = action.payload.access_token;
		});
		builder.addCase(login.rejected, (state, action) => {
			state.loginErrorMessage = action.error.message;
		});
		builder.addCase(getProfile.fulfilled, (state, action) => {
			state.profile = action.payload;
		});
	}
});

export default userSlice.reducer;
export const userActions = userSlice.actions;
