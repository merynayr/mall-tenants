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

const router = createBrowserRouter([
	{
		path: '/',
		element: <RequireAuth><Layout /></RequireAuth>,
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
			},
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
		element:  <Page404 />
	}
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
	<React.StrictMode>
		<Provider store={store}>
			<RouterProvider router={router} />
		</Provider>
	</React.StrictMode>
);
