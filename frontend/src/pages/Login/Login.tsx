import { FormEvent, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import styles from './Login.module.css';
import Button from '@/components/Button/Button';
import Headling from '@/components/Headling/Headling';
import Input from '@/components/Input/Input';
import { AppDispath, RootState } from '@/store/store';
import { getProfile, login, userActions } from '@/store/user.slice';
import { toast } from 'react-toastify';

export type LoginForm = {
	email: {
		value: string;
	};
	password: {
		value: string;
	};
}

export function Login() {
	const navigate = useNavigate();
	const dispatch = useDispatch<AppDispath>();
	const { jwt, loginErrorMessage } = useSelector((s: RootState) => s.user);

	useEffect(() => {
		if (loginErrorMessage) {
			toast.error(loginErrorMessage);
		}
	}, [loginErrorMessage]);

	useEffect(() => {
		if (jwt) {
			dispatch(getProfile()); 
			navigate('/');
		}
	}, [dispatch, jwt, navigate]);

	const submit = async (e: FormEvent) => {
		e.preventDefault();
		dispatch(userActions.clearLoginError());
		const target = e.target as typeof e.target & LoginForm;
		const { email, password } = target;
		await sendLogin(email.value, password.value);
	};

	const sendLogin = async (email: string, password: string) => {
		dispatch(login({ email, password }));
	};

	return <div className={styles['login']}>
		<Headling>Вход</Headling>
		<form className={styles['form']} onSubmit={submit}>
			<div className={styles['field']}>
				<label htmlFor="email">Ваш email</label>
				<Input id="email" name='email' placeholder='Email' />
			</div>
			<div className={styles['field']}>
				<label htmlFor="password">Ваш пароль</label>
				<Input id="password" name='password' type="password" placeholder='Пароль' />
			</div>
			<Button appearence="big">Вход</Button>
		</form>
	</div>;
}