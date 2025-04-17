import React, { lazy } from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { RouterProvider, createBrowserRouter  } from 'react-router-dom';
import { Layout } from '@/layout/Menu/Layout';
import Clients from '@/pages/Clients/Clients';
import Payments from '@/pages/Payments/Payments';
import Premises from '@/pages/Premises/Premises';

const Menu = lazy(() => import('@/pages/Menu/Menu'));

const router = createBrowserRouter([
	{
		path: '/',
		element: <Layout />,
		children: [
			{
				path: '/',
				element: <Menu />
			},
			{
				path: '/premises',
				element: <Premises />
			},
			{
				path: '/clients',
				element: <Clients />
			},
			{
				path: '/payments',
				element: <Payments />
			}
		]
	},
	{
		path: '*',
		element: <Layout />
	}
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
	<React.StrictMode>		
		<RouterProvider router={router} />
	</React.StrictMode>
);