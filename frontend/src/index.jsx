import * as React from 'react';
import '@ant-design/v5-patch-for-react-19';
import { createRoot } from 'react-dom/client';
import { RouterProvider } from 'react-router';
import { Provider } from 'jotai';
import { RefreshTokenUtil } from 'service/helper/request_util';
import 'service/styles/main.css';
import router from './router';

window.refreshTokenUtil = new RefreshTokenUtil();
window.socConn = null

createRoot(document.getElementById('root')).render(
    <Provider>
        <RouterProvider router={router} />
    </Provider>
);
