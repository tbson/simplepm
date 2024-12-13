import * as React from 'react';
import { useState, useEffect } from 'react';
import { t } from 'ttag';
import { Modal, Button } from 'antd';
import Util from 'service/helper/util';
import Form from './form';
import { getMessages, TOGGLE_DIALOG_EVENT } from '../config';

export class Service {
    static get toggleEvent() {
        return TOGGLE_DIALOG_EVENT;
    }

    static toggle(open = true, data = null) {
        Util.event.dispatch(Service.toggleEvent, { open, data });
    }
}

/**
 * TaskFieldOptionDialog.
 *
 * @param {Object} props
 * @param {function} props.onChange - (data: Dict, id: number) => void
 */
export default function TaskFieldOptionDialog({ onChange, onDelete }) {
    const [data, setData] = useState({});
    const [open, setOpen] = useState(false);
    const messages = getMessages();

    const handleToggle = ({ detail: { open, data } }) => {
        setOpen(open);
        setData(data || {});
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
            title={Util.getDialogTitle(data.id, messages)}
            footer={(_, { OkBtn, CancelBtn }) => (
                <div style={{ display: 'flex' }}>
                    <div style={{ width: 40 }}>
                        {data.id ? (
                            <Button
                                danger
                                onClick={() => {
                                    onDelete(data.id);
                                    Service.toggle(false);
                                }}
                            >{t`Delete`}</Button>
                        ) : null}
                    </div>
                    <div style={{ flex: 1 }}>
                        <CancelBtn />
                        &nbsp; &nbsp;
                        <OkBtn />
                    </div>
                </div>
            )}
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

TaskFieldOptionDialog.displayName = 'TaskFieldOptionDialog';
TaskFieldOptionDialog.toggle = Service.toggle;
