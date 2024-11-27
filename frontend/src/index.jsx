import * as React from 'react';
import { createRoot } from 'react-dom/client';
import { RouterProvider } from 'react-router-dom';
import { Provider } from 'jotai'
import 'service/styles/main.css';
import router from './router';

createRoot(document.getElementById('root')).render(
    <Provider>
        <RouterProvider router={router} />
    </Provider>
);
