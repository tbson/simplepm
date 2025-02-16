import * as React from 'react';
import { useEffect } from 'react';
import { useAtom } from 'jotai';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { projectOptionSt } from './state';
import { urls, getMessages } from './config';
import Table from './table';

export default function Project() {
    const [projectOption, setProjectOption] = useAtom(projectOptionSt);
    useEffect(() => {
        if (!projectOption.loaded) {
            getOption();
        }
    }, []);

    function getOption() {
        RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setProjectOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setProjectOption((prev) => ({ ...prev, loaded: true }));
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
