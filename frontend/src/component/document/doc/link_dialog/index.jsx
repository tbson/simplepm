import React, { useState, useEffect } from 'react';
import { t } from 'ttag';
import { Modal } from 'antd';
import Util from 'service/helper/util';
import Form from './form';
import { getMessages, TOGGLE_LINK_DIALOG_EVENT } from '../config';

export class Service {
    static get toggleEvent() {
        return TOGGLE_LINK_DIALOG_EVENT;
    }

    static toggle(open = true) {
        Util.event.dispatch(Service.toggleEvent, { open });
    }
}

/**
 * DocLinkDialog.
 *
 * @param {Object} props
 * @param {function} props.onChange - (data: Dict) => void
 */
export default function DocLinkDialog({ onChange }) {
    const [open, setOpen] = useState(false);
    const messages = getMessages();

    const handleToggle = ({ detail: { open } }) => {
        if (!open) {
            return setOpen(false);
        }
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
            destroyOnClose
            open={open}
            okButtonProps={{ form: Form.formName, key: 'submit', htmlType: 'submit' }}
            okText={t`Save`}
            onCancel={() => Service.toggle(false)}
            cancelText={t`Cancel`}
            title={Util.getDialogTitle(null, messages)}
        >
            <Form
                onChange={(data) => {
                    setOpen(false);
                    onChange(data);
                }}
            />
        </Modal>
    );
}

DocLinkDialog.displayName = 'DocLinkDialog';
DocLinkDialog.toggle = Service.toggle;
