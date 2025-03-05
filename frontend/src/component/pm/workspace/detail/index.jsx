import React, { useEffect, useState } from 'react';
import { t } from 'ttag';
import { useAtom } from 'jotai';
import { useParams } from 'react-router';
import { App, Divider, Button } from 'antd';
import { EditOutlined } from '@ant-design/icons';
import PageHeading from 'component/common/page_heading';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import { workspaceOptionSt } from '../state';
import { urls, getMessages } from '../config';
import Summary from './summary';
import Dialog from '../dialog';

export default function Workspace() {
    const { notification } = App.useApp();
    const { workspace_id } = useParams();
    const [item, setItem] = useState({});
    const [workspaceOption, setWorkspaceOption] = useAtom(workspaceOptionSt);
    useEffect(() => {
        if (!workspaceOption.loaded) {
            getOption();
        }
        getItem(workspace_id);
    }, []);

    const getOption = () => {
        return RequestUtil.apiCall(urls.option)
            .then((resp) => {
                setWorkspaceOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setWorkspaceOption((prev) => ({ ...prev, loaded: true }));
            });
    };

    const getItem = (workspace_id) => {
        Util.toggleGlobalLoading();
        return RequestUtil.apiCall(`${urls.crud}${workspace_id}`)
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
                    onClick={() => Dialog.toggle(true, workspace_id)}
                >
                    {t`Update workspace`}
                </Button>
            </div>
            <br />
            <Dialog onChange={onChange} />
        </>
    );
}

Workspace.displayName = 'WorkspaceDetail';
