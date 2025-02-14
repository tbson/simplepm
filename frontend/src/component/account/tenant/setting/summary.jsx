import * as React from 'react';
import { t } from 'ttag';
import { Button } from 'antd';
import { GithubOutlined } from '@ant-design/icons';
import RequestUtil from 'service/helper/request_util';
import Img from 'component/common/display/img';
import { githubUrls } from '../config';

export default function TenantSettingSummary({ data }) {
    const getGithubInstallUrl = () => {
        return RequestUtil.apiCall(githubUrls.installUrl).then((resp) => {
            console.log(resp.data.url);
            // window.location.href = resp.data.url;
        });
    };

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
                        <strong>{t`Code`}</strong>
                    </td>
                    <td span={18}>{data.uid}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`Name`}</strong>
                    </td>
                    <td span={18}>{data.title}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`Github account`}</strong>
                    </td>
                    <td span={18}>
                        <Button
                            icon={<GithubOutlined />}
                            onClick={() => {
                                getGithubInstallUrl();
                            }}
                        >
                            Connect to Github account
                        </Button>
                    </td>
                </tr>
            </tbody>
        </table>
    );
}
TenantSettingSummary.displayName = 'TenantSettingSummary';
