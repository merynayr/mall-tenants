import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { Provider } from 'react-redux';
import { RouterProvider, createBrowserRouter  } from 'react-router-dom';
import { RequireAuth } from '@/helpers/RequireAuth';
import { AuthLayout } from '@/layout/Auth/AuthLayout';
import { Layout } from '@/layout/Menu/Layout';
import Page404 from '@/pages/404/404';
import PageClients from '@/pages/Clients/Clients';
import { Login } from '@/pages/Login/Login';
import Payments from '@/pages/Payments/Payments';
import PagePremises from '@/pages/Premise/Premise';
import PremiseInfo from '@/pages/PremiseInfo/PremiseInfo';
import PremisesMapperPage from '@/pages/PremisesMapper/PremisesMapperPage';
import PageRents from '@/pages/Rents/Rents';
import { store } from '@/store/store';
import PageMyRents from './pages/MyRents/MyRents';
import { PageApplications } from './pages/Applications/Applications';
import PageMyPayments from './pages/MyPayments/MyPayments';
import { ToastContainer } from 'react-toastify';

const router = createBrowserRouter([
	{
		path: '/',
		element: <Layout />, 
		children: [
			{
				path: '/',
				element: <PremisesMapperPage />
			},
			{
				path: '/premises',
				element: <PagePremises />
			},
			{
				path: '/premise/:id',
				element: <PremiseInfo />
			}
		]
	},
	{
		path: '/',
		element: (
			<RequireAuth>
				<Layout />
			</RequireAuth>
		),
		children: [
			{
				path: '/clients',
				element: <PageClients />
			},
			{
				path: '/payments',
				element: <Payments />
			},
			{
				path: '/rents',
				element: <PageRents />
			},
			{
				path: '/my-rents',
				element: <PageMyRents />
			},
			{
				path: '/my-payments',
				element: <PageMyPayments />
			},
			{
				path: '/applications',
				element: <PageApplications />
			}
		]
	},
	{
		path: '/auth',
		element: <AuthLayout />,
		children: [
			{
				path: 'login',
				element: <Login />
			}
		]
	},
	{
		path: '*',
		element: <Page404 />
	}
]);


ReactDOM.createRoot(document.getElementById('root')!).render(
	<React.StrictMode>
		<Provider store={store}>
			<RouterProvider router={router} />
			<ToastContainer position="top-right" autoClose={3000} />
		</Provider>
	</React.StrictMode>
);
