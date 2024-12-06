import * as React from 'react';
import { useEffect } from 'react';
import { useAtom } from 'jotai';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { projectOptionSt } from './state';
import { urls, getMessages } from './config';
import Table from './table';

export default function Project() {
    const [option, setOption] = useAtom(projectOptionSt);
    useEffect(() => {
        if (!option.loaded) getOption();
    }, []);

    function getOption() {
        RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setOption((prev) => ({ ...prev, loaded: true }));
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

Project.displayName = 'Project';
