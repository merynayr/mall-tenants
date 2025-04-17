import cn from 'classnames';
import { useState } from 'react';
import { NavLink, Outlet } from 'react-router-dom';
import styles from './Layout.module.css';
import Button from '@/components/Button/Button';

export function Layout() {
	const [isSidebarOpen, setIsSidebarOpen] = useState(true);

	const toggleSidebar = () => {
		setIsSidebarOpen(prev => !prev);
	};

	return <>
		<header className={styles['header']}>
			<div className={styles['header-left']}>
				<div className={styles['company-name']}>🏢 ООО "Твой Бизнес"</div>

			</div>
			<div className={styles['header-right']}>
				<button className={styles['settings-button']}>⚙️</button>
			</div>
		</header>
		<div className={styles['layout']}>
			<div className={cn(styles.sidebar, { [styles['sidebar-closed']]: !isSidebarOpen })}>
				<div className={styles['user']}>
					{isSidebarOpen && (
						<>
							<div>
								<div className={styles['name']}>Dima</div>
								<div className={styles['email']}>1@mail.ru</div>
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

					<NavLink to="/premises" className={({ isActive }) => cn(styles['link'], { [styles.active]: isActive })}>
						 <img src="/premise-icon.svg" alt="Иконка помещений" className={styles['icon']} />
						{isSidebarOpen && 'Помещения'}
					</NavLink>

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
				</div>
				<Button className={styles['exit']}>
					<img src="/exit-icon.svg" alt="Иконка выхода" className={styles['icon']} />
					{isSidebarOpen && 'Выход'}
				</Button>
			</div>
			<div className={styles['content']}>
				<Outlet />
			</div>
		</div>
	</>;
}
