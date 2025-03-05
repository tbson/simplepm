import React from 'react';
import PageHeading from 'component/common/page_heading';
import { getMessages } from './config';
import Table from './table';

export default function Workspace() {
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

Workspace.displayName = 'Workspace';
