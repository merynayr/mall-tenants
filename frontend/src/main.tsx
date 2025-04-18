import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { RouterProvider, createBrowserRouter  } from 'react-router-dom';
import { Layout } from '@/layout/Menu/Layout';
import Clients from '@/pages/Clients/Clients';
import DrawTesterPage from '@/pages/Menu/Menu';
import Payments from '@/pages/Payments/Payments';
import PagePremises from '@/pages/Premise/Premise';
import PremiseInfo from '@/pages/PremiseInfo/PremiseInfo';

// const Menu = lazy(() => import('@/pages/Menu/Menu'));

const router = createBrowserRouter([
	{
		path: '/',
		element: <Layout />,
		children: [
			{
				path: '/',
				element: <DrawTesterPage />
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
				element: <Clients />
			},
			{
				path: '/payments',
				element: <Payments />
			},
			{
				path: '/rents',
				element: <></>
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