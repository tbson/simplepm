import React, { useEffect, useState } from 'react';
import { t } from 'ttag';
import { useAtom } from 'jotai';
import { useParams } from 'react-router';
import { App, Divider, Button } from 'antd';
import { EditOutlined } from '@ant-design/icons';
import PageHeading from 'component/common/page_heading';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import { tenantOptionSt } from '../state';
import { urls, getMessages } from '../config';
import Summary from './summary';
import Dialog from '../dialog';

export default function TenantDetail() {
    const { notification } = App.useApp();
    const { tenant_id } = useParams();
    const [item, setItem] = useState({});
    const [tenantOption, setTenantOption] = useAtom(tenantOptionSt);
    useEffect(() => {
        if (!tenantOption.loaded) {
            getOption();
        }
        getItem(tenant_id);
    }, []);

    const getOption = () => {
        return RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setTenantOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setTenantOption((prev) => ({ ...prev, loaded: true }));
            });
    };

    const getItem = (tenant_id) => {
        Util.toggleGlobalLoading();
        return RequestUtil.apiCall(`${urls.crud}${tenant_id || 0}`)
            .then((resp) => {
                setItem(resp.data);
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                Util.toggleGlobalLoading(false);
            });
    };

    const onChange = (data, _id) => {
        setItem(data);
    };

    const messages = getMessages();
    return (
        <>
            <PageHeading>
                <>{messages.heading}</>
            </PageHeading>
            <Summary data={item} />
            <Divider />
            <div className="right">
                <Button
                    htmlType="button"
                    type="primary"
                    icon={<EditOutlined />}
                    onClick={() => Dialog.toggle(true, tenant_id)}
                >
                    {t`Update tenant`}
                </Button>
            </div>
            <br />
            <Dialog onChange={onChange} />
        </>
    );
}

TenantDetail.displayName = 'TenantDetail';
