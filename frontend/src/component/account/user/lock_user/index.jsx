import React, { useState, useEffect } from 'react';
import { t } from 'ttag';
import { App, Modal } from 'antd';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import Form from './form';
import { urls, TOGGLE_LOCK_DIALOG_EVENT } from '../config';

export class Service {
    static get toggleEvent() {
        return TOGGLE_LOCK_DIALOG_EVENT;
    }

    static toggle(open = true, id = 0) {
        Util.event.dispatch(Service.toggleEvent, { open, id });
    }
}

/**
 * UserDialog.
 *
 * @param {Object} props
 * @param {function} props.onChange - (data: Dict, id: number) => void
 */
export default function UserDialog({ onChange }) {
    const { notification } = App.useApp();
    const [data, setData] = useState({});
    const [open, setOpen] = useState(false);

    const handleToggle = ({ detail: { open, id } }) => {
        if (!open) {
            return setOpen(false);
        }
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
            setData({});
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
            destroyOnHidden
            open={open}
            okButtonProps={{ form: Form.formName, key: 'submit', htmlType: 'submit' }}
            okText={t`Save`}
            onCancel={() => Service.toggle(false)}
            cancelText={t`Cancel`}
            title={`Update lock status of ${data.email}`}
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

UserDialog.displayName = 'UserDialog';
UserDialog.toggle = Service.toggle;
