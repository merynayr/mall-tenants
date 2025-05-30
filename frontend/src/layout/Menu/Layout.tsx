import cn from 'classnames';
import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { NavLink, Outlet, useNavigate } from 'react-router-dom';
import styles from './Layout.module.css';
import Button from '@/components/Button/Button';
import { useHasRole } from '@/hooks/Role';
import { AppDispath, RootState } from '@/store/store';
import {  userActions } from '@/store/user.slice';


export function Layout() {
	const [isSidebarOpen, setIsSidebarOpen] = useState(true);
	const navigate = useNavigate();
	const dispatch = useDispatch<AppDispath>();
	const { profile } = useSelector((state: RootState) => state.user);
	const hasRole = useHasRole('client', 'moderator', 'director');

	const toggleSidebar = () => {
		setIsSidebarOpen(prev => !prev);
	};

	const logout = () => {
		document.cookie = 'access_token=; Max-Age=-1; path=/';
		document.cookie = 'refresh_token=; Max-Age=-1; path=/';
		dispatch(userActions.logout());
		navigate('/auth/login');
	};

	return <>
		<header className={styles['header']}>
			<div className={styles['header-left']}>
				<div className={styles['company-name']}>
					<img src="/favicon.png" alt="Rentify" className={styles['favicon']} />
  ООО "Rentify"
				</div>

			</div>
		</header>
		<div className={styles['layout']}>
			<div className={cn(styles.sidebar, { [styles['sidebar-closed']]: !isSidebarOpen })}>
				<div className={styles['user']}>
					{isSidebarOpen && (
						<>
							<div>
								<div className={styles['email']}>{profile?.email}</div>
							</div>
						</>
					)}
					<button className={styles['settings-button']} onClick={toggleSidebar}>
						<img src="/sidebar-icon.svg" alt="Иконка главной" className={styles['icon']} />
					</button>
				</div>
				<div className={styles['menu']}>
					<NavLink to="/" className={({ isActive }) => cn(styles['link'], { [styles.active]: isActive })}>
						<img src="/menu-icon.svg" alt="Иконка главной" className={styles['icon']} />
						{isSidebarOpen && 'Основной экран'}
					</NavLink>

						{useHasRole('moderator') && (
						<>
							<NavLink to="/applications" className={({ isActive }) => cn(styles['link'], { [styles.active]: isActive })}>
								<img src="/application-icon.svg" alt="Иконка заявок" className={styles['icon']} />
								{isSidebarOpen && 'Заявки'}
							</NavLink>
						</>
						)}
						
						<NavLink to="/premises" className={({ isActive }) => cn(styles['link'], { [styles.active]: isActive })}>
								<img src="/premise-icon.svg" alt="Иконка помещений" className={styles['icon']} />
								{isSidebarOpen && 'Помещения'}
							</NavLink>

						{useHasRole('moderator', 'director') && (
						<>
							<NavLink to="/clients" className={({ isActive }) => cn(styles['link'], { [styles.active]: isActive })}>
								<img src="/client-icon.svg" alt="Иконка клиентов" className={styles['icon']} />
								{isSidebarOpen && 'Клиенты'}
							</NavLink>
							<NavLink to="/payments" className={({ isActive }) => cn(styles['link'], { [styles.active]: isActive })}>
								<img src="/payment-icon.svg" alt="Иконка оплат" className={styles['icon']} />
								{isSidebarOpen && 'Оплаты'}
							</NavLink>
							<NavLink to="/rents" className={({ isActive }) => cn(styles['link'], { [styles.active]: isActive })}>
								<img src="/rent-icon.svg" alt="Иконка договоров" className={styles['icon']} />
								{isSidebarOpen && 'Договора'}
							</NavLink>
						</>
					)}

						{useHasRole('client') && (
						<>
							<NavLink to="/my-rents" className={({ isActive }) => cn(styles['link'], { [styles.active]: isActive })}>
								<img src="/rent-icon.svg" alt="Иконка договоров" className={styles['icon']} />
								{isSidebarOpen && 'Мои аренды'}
							</NavLink>
							<NavLink to="/my-payments" className={({ isActive }) => cn(styles['link'], { [styles.active]: isActive })}>
								<img src="/payment-icon.svg" alt="Иконка оплат" className={styles['icon']} />
								{isSidebarOpen && 'Мои платежи'}
							</NavLink>
						</>
					)}
				</div>
				<Button className={styles['exit']} onClick={logout}>
					<img src="/exit-icon.svg" alt="Иконка выхода" className={styles['icon']} />
					{isSidebarOpen && (hasRole ? 'Выход' : 'Вход')}
				</Button>
			</div>
			<div className={styles['content']}>
				<Outlet />
			</div>
		</div>
	</>;
}
