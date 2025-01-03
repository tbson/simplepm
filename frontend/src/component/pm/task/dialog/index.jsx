import * as React from 'react';
import { useState, useEffect } from 'react';
import { t } from 'ttag';
import { Modal } from 'antd';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import Form from './form';
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
};

/**
 * TaskDialog.
 *
 * @param {Object} props
 * @param {function} props.onChange - (data: Dict, id: number) => void
 */
export default function TaskDialog({ onChange }) {
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

    return (
        <Modal
            keyboard={false}
            maskClosable={false}
            destroyOnClose
            open={open}
            okButtonProps={{ form: Form.formName, key: 'submit', htmlType: 'submit' }}
            okText={t`Save`}
            onCancel={() => Service.toggle(false)}
            cancelText={t`Cancel`}
            title={Util.getDialogTitle(id, messages)}
        >
            <Form
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
