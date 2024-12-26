import * as React from 'react';
import { useEffect } from 'react';
import { useLocale } from 'ttag';
import { App, ConfigProvider } from 'antd';
import { Outlet } from 'react-router-dom';
import Spinner from 'component/common/spinner';
import Util from 'service/helper/util';
import LocaleUtil from 'service/helper/locale_util';

Util.responseIntercept();

const themeConfig = {
    components: { Menu: { itemHeight: 34 } },
    token: {
        fontFamily: 'Montserrat',
        colorPrimary: '#255891',
        colorLink: '#255891',
        borderRadius: 3,
    }
};

export default function MainApp() {
    useEffect(() => {
        LocaleUtil.fetchLocales().then(() => {
            useLocale(LocaleUtil.getLocale());
        }).catch((err) => {
            console.error(err);
        });
    }, []);

    return (
        <div>
            <ConfigProvider theme={themeConfig}>
                <App>
                    <Spinner />
                    <Outlet />
                </App>
            </ConfigProvider>
        </div>
    );
}
