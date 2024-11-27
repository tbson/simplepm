import * as React from 'react';
import { addLocale, useLocale } from 'ttag';
import { App, ConfigProvider } from 'antd';
import { Outlet } from 'react-router-dom';
import Spinner from 'component/common/spinner';
import Util from 'service/helper/util';
import LocaleUtil from 'service/helper/locale_util';
import vi from 'src/locale/vi.po.json';
import en from 'src/locale/en.po.json';
const langs = { vi, en };

Util.responseIntercept();

const themeConfig = {
    components: { Menu: { itemHeight: 34 } },
    token: {
        fontFamily: 'Montserrat',
        colorPrimary: '#006576',
        colorLink: '#006576',
        borderRadius: 3,
    }
};

export default function MainApp() {
    LocaleUtil.getSupportedLocales().forEach((locale) => {
        addLocale(locale, langs[locale]);
    });
    useLocale(LocaleUtil.getLocale());

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
