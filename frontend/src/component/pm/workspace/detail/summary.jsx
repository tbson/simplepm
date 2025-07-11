import React from 'react';
import { useAtomValue } from 'jotai';
import { t } from 'ttag';
import Img from 'component/common/display/img';
import { workspaceDictSt } from 'component/account/workspace/state';

export default function WorkspaceSummary({ data }) {
    const workspaceDict = useAtomValue(workspaceDictSt);
    return (
        <table className="styled-table">
            <tbody>
                <tr>
                    <td span={6}>
                        <strong>{t`Avatar`}</strong>
                    </td>
                    <td span={18}>
                        <Img src={data.avatar} width={150} height={150} />
                    </td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`UID`}</strong>
                    </td>
                    <td span={18}>{data.uid}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`Title`}</strong>
                    </td>
                    <td span={18}>{data.title}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`Auth client`}</strong>
                    </td>
                    <td span={18}>
                        {workspaceDict.auth_client[data.auth_client_id] || ''}
                    </td>
                </tr>
            </tbody>
        </table>
    );
}
WorkspaceSummary.displayName = 'WorkspaceSummary';
