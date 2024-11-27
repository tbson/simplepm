import * as React from 'react';
import PageHeading from 'component/common/page_heading';
import { getMessages } from './config';
import Table from './table';

export default function AuthClient() {
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

AuthClient.displayName = 'AuthClient';
