import React, { useState, useEffect } from 'react';
import { t } from 'ttag';
import { Modal } from 'antd';
import Util from 'service/helper/util';
import Form from './form';
import { TOGGLE_DIALOG_EVENT } from '../config';

export class Service {
    static get toggleEvent() {
        return TOGGLE_DIALOG_EVENT;
    }

    static toggle(open = true) {
        Util.event.dispatch(Service.toggleEvent, { open });
    }
}

/**
 * ResetPwdDialog.
 *
 * @param {Object} props
 * @param {function} props.onChange - (data: Dict) => void
 */
export default function ResetPwdDialog() {
    const [data, setData] = useState({});
    const [open, setOpen] = useState(false);

    const handleToggle = ({ detail: { open } }) => {
        if (!open) {
            return setOpen(false);
        }
        setData({});
        setOpen(true);
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
            title={t`Reset Password`}
        >
            <Form
                data={data}
                onChange={() => {
                    setOpen(false);
                }}
            />
        </Modal>
    );
}

ResetPwdDialog.displayName = 'ResetPwdDialog';
ResetPwdDialog.toggle = Service.toggle;
