import axios, { AxiosError, InternalAxiosRequestConfig } from 'axios';
import { loadState, saveState } from '@/store/storage';
import { JWT_PERSISTENT_STATE, UserPersistentState } from '@/store/user.slice';

interface CustomAxiosRequestConfig extends InternalAxiosRequestConfig {
  _retry?: boolean;
}

export const PREFIX = 'http://localhost:8080';

const instance = axios.create({
	baseURL: PREFIX,
	withCredentials: true
});

// === Request Interceptor ===
instance.interceptors.request.use((config) => {
	const state = loadState<UserPersistentState>(JWT_PERSISTENT_STATE);
	const token = state?.jwt;

	if (token && config.headers) {
		config.headers.Authorization = `Bearer ${token}`;
	}

	return config;
}, (error) => {
	return Promise.reject(error);
});

// === Refresh Token Logic ===
let isRefreshing = false;
let failedQueue: {
  resolve: (token: string) => void;
  reject: (err: AxiosError | unknown) => void;
}[] = [];

const processQueue = (error: AxiosError | null, token: string | null = null) => {
	failedQueue.forEach(prom => {
		if (token) {
			prom.resolve(token);
		} else {
			prom.reject(error);
		}
	});
	failedQueue = [];
};

async function refreshToken(): Promise<string> {
	const response = await axios.post<{ access_token: string, refresh_token: string }>(
		`${PREFIX}/auth/refresh`,
		{},
		{ withCredentials: true }
	);

	const newToken = response.data.access_token;

	const prevState = loadState(JWT_PERSISTENT_STATE) || {};
	saveState(
		{ ...prevState, jwt: newToken },
		JWT_PERSISTENT_STATE
	);
	return newToken;
}

function logoutUser() {
	saveState(null, JWT_PERSISTENT_STATE);
	window.location.href = '/auth/login';
}

// === Response Interceptor ===
instance.interceptors.response.use(
	(response) => response,
	async (error: AxiosError) => {
		const originalRequest = error.config as CustomAxiosRequestConfig;

		if (error.response?.status === 401  && !originalRequest._retry) {
			if (isRefreshing) {
				return new Promise((resolve, reject) => {
					failedQueue.push({
						resolve: (token: string) => {
							if (originalRequest.headers) {
								originalRequest.headers.Authorization = `Bearer ${token}`;
							}
							resolve(instance(originalRequest));
						},
						reject: (err) => reject(err)
					});
				});
			}

			originalRequest._retry = true;
			isRefreshing = true;

			try {
				const newToken = await refreshToken();
				processQueue(null, newToken);

				if (originalRequest.headers) {
					originalRequest.headers.Authorization = `Bearer ${newToken}`;
				}

				return instance(originalRequest);
			} catch (refreshError) {
				if (refreshError instanceof AxiosError) {
					processQueue(refreshError, null);
				}
				console.error('Не удалось обновить токен:', refreshError);
				logoutUser();
				return Promise.reject(refreshError);
			} finally {
				isRefreshing = false;
			}
		}

		return Promise.reject(error);
	}
);

export default instance;
