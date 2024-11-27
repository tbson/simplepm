import * as React from 'react';
import { useEffect } from 'react';
import { useAtom } from 'jotai';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { tenantOptionSt } from './state';
import { urls, getMessages } from './config';
import Table from './table';

export default function Tenant() {
    const [tenantOption, setTenantOption] = useAtom(tenantOptionSt);
    useEffect(() => {
        if (!tenantOption.loaded) {
            getOption();
        }
    }, []);

    const getOption = () => {
        RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setTenantOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setTenantOption((prev) => ({ ...prev, loaded: true }));
            });
    };

    const messages = getMessages();
    return (
        <>
            <PageHeading>
                <>{messages.heading}</>
            </PageHeading>
            <Table />
        </>
    );
}

Tenant.displayName = 'Tenant';
