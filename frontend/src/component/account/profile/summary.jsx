import * as React from 'react';
import { t } from 'ttag';
import Img from 'component/common/display/img';

export default function ProfileSummary(data) {
    return (
        <table className="styled-table">
            <tbody>
                <tr>
                    <td span={6}>
                        <strong>{t`Avatar`}</strong>
                    </td>
                    <td span={18}>
                        <Img
                            src={data.avatar}
                            width={150}
                            height={150}
                        />
                    </td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`Email`}</strong>
                    </td>
                    <td span={18}>{data.email}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`Mobile`}</strong>
                    </td>
                    <td span={18}>{data.mobile}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`First name`}</strong>
                    </td>
                    <td span={18}>{data.first_name}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`Last name`}</strong>
                    </td>
                    <td span={18}>{data.last_name}</td>
                </tr>
            </tbody>
        </table>
    );
}
ProfileSummary.displayName = 'ProfileSummary';
