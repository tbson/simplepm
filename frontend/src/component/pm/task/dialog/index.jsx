import * as React from 'react';
import { t } from 'ttag';
import { useState, useEffect } from 'react';
import { App, Modal, Button } from 'antd';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import TaskForm from './form';
import { urls, getMessages, TOGGLE_DIALOG_EVENT } from '../config';

export class Service {
    static get toggleEvent() {
        return TOGGLE_DIALOG_EVENT;
    }

    static toggle(open = true, id = 0, status = 0) {
        Util.event.dispatch(Service.toggleEvent, { open, id, status });
    }
}

const emptyData = {
    id: 0,
    title: '',
    description: '',
    feature_id: null,
    task_fields: [],
    task_users: [],
};

/**
 * TaskDialog.
 *
 * @param {Object} props
 * @param {function} props.onChange - (data: Dict, id: number) => void
 */
export default function TaskDialog({ projectId, onChange, onDelete }) {
    const { notification } = App.useApp();
    const [data, setData] = useState(emptyData);
    const [open, setOpen] = useState(false);
    const [id, setId] = useState(0);
    const messages = getMessages();

    const handleToggle = ({ detail: { open, id, status } }) => {
        if (!open) {
            return setOpen(false);
        }
        setId(id);
        if (id) {
            Util.toggleGlobalLoading();
            RequestUtil.apiCall(`${urls.crud}${id}`)
                .then((resp) => {
                    setData(resp.data);
                    setOpen(true);
                })
                .catch(RequestUtil.displayError(notification))
                .finally(() => Util.toggleGlobalLoading(false));
        } else {
            const data = { ...emptyData, status };
            setData(data);
            setOpen(true);
        }
    };

    useEffect(() => {
        Util.event.listen(Service.toggleEvent, handleToggle);
        return () => {
            Util.event.remove(Service.toggleEvent, handleToggle);
        };
    }, []);

    const handleDelete = (id) => {
        console.log('delete', id);
    };

    return (
        <Modal
            keyboard={false}
            maskClosable={false}
            destroyOnClose
            open={open}
            okButtonProps={{
                form: TaskForm.formName,
                key: 'submit',
                htmlType: 'submit'
            }}
            okText={t`Save`}
            onCancel={() => Service.toggle(false)}
            cancelText={t`Cancel`}
            title={Util.getDialogTitle(id, messages)}
            footer={(_, {}) => (
                <div style={{ display: 'flex' }}>
                    <div style={{ width: 40 }}>
                        <Button onClick={() => Service.toggle(false)}>Cancel</Button>
                    </div>
                    <div style={{ flex: 1 }} className="right">
                        {id ? (
                            <Button
                                danger
                                onClick={() => onDelete(id)}
                            >{t`Delete`}</Button>
                        ) : null}
                        &nbsp; &nbsp;
                        <Button
                            type="primary"
                            form={TaskForm.formName}
                            key="submit"
                            htmlType="submit"
                        >
                            {id ? 'Update' : 'Create'}
                        </Button>
                    </div>
                </div>
            )}
        >
            <TaskForm
                projectId={projectId}
                data={data}
                onChange={(data, id) => {
                    setOpen(false);
                    onChange(data, id);
                }}
            />
        </Modal>
    );
}

TaskDialog.displayName = 'TaskDialog';
TaskDialog.toggle = Service.toggle;
