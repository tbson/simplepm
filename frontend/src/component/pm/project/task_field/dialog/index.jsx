import * as React from 'react';
import { t } from 'ttag';
import { App, Button } from 'antd';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import Form from './form';
import { urls, getMessages } from '../config';

/**
 * TaskFieldDialog.
 *
 * @param {Object} props
 * @param {function} props.onChange - (data: Dict, id: number) => void
 */
export default function TaskFieldDialog({ id, data, toggle, onChange, onDelete }) {
    const { notification } = App.useApp();
    const messages = getMessages();

    const handleDelete = (id) => {
        const r = window.confirm(messages.deleteOne);
        if (!r) {
            return;
        }
        Util.toggleGlobalLoading(true);
        RequestUtil.apiCall(`${urls.crud}${id}`, {}, 'delete')
            .then(() => {
                onDelete(id);
                toggle(false);
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => Util.toggleGlobalLoading(false));
    };
    return (
        <div>
            <Form
                data={data}
                onChange={(data, id) => {
                    onChange(data, id);
                    toggle(false);
                }}
            />
            <div style={{ display: 'flex' }}>
                <div style={{ width: 40 }}>
                    <Button onClick={() => toggle(false)}>Cancel</Button>
                </div>
                <div style={{ flex: 1 }} className="right">
                    {id ? (
                        <Button
                            danger
                            onClick={() => handleDelete(id)}
                        >{t`Delete`}</Button>
                    ) : null}
                    &nbsp; &nbsp;
                    <Button
                        type="primary"
                        form={Form.formName}
                        key="submit"
                        htmlType="submit"
                    >
                        {id ? 'Update' : 'Create'}
                    </Button>
                </div>
            </div>
        </div>
    );
}

TaskFieldDialog.displayName = 'TaskFieldDialog';
