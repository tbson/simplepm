import * as React from 'react';
import { useEffect } from 'react';
import { useAtom } from 'jotai';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { roleOptionSt } from './state';
import { urls, getMessages } from './config';
import Table from './table';

export default function Role() {
    const [roleOption, setRoleOption] = useAtom(roleOptionSt);
    useEffect(() => {
        if (!roleOption.loaded) {
            getOption();
        }
    }, []);

    const getOption = () => {
        RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setRoleOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setRoleOption((prev) => ({ ...prev, loaded: true }));
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

Role.displayName = 'Role';
