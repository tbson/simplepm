import * as React from 'react';
import { useEffect } from 'react';
import { useAtom } from 'jotai';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { userOptionSt } from './state';
import { urls, getMessages } from './config';
import Table from './table';

export default function User() {
    const [userOption, setUserOption] = useAtom(userOptionSt);
    useEffect(() => {
        if (!userOption.loaded) getOption();
    }, []);

    function getOption() {
        RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setUserOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setUserOption((prev) => ({ ...prev, loaded: true }));
            });
    }

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

User.displayName = 'User';
