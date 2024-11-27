import * as React from 'react';
import { useEffect } from 'react';
import { useAtom } from 'jotai';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { variableOptionSt } from './state';
import { urls, getMessages } from './config';
import Table from './table';

export default function Variable() {
    const [variableOption, setVariableOption] = useAtom(variableOptionSt);
    useEffect(() => {
        if (!variableOption.loaded) getOption();
    }, []);

    function getOption() {
        RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setVariableOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setVariableOption((prev) => ({ ...prev, loaded: true }));
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

Variable.displayName = 'Variable';
