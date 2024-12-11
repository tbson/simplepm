import React from 'react';
import { useState, useEffect } from 'react';
import { useSetAtom } from 'jotai';
import Util from 'service/helper/util';
import { Drawer } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { IconButton } from 'component/common/table/buttons';
import { projectIdSt } from './state';
import TaskFieldTable from './table';
import TaskFieldDialog from './dialog';
import { getMessages, TOGGLE_TASK_FIELD_EVENT } from './config';

export class Service {
    static get toggleEvent() {
        return TOGGLE_TASK_FIELD_EVENT;
    }

    static toggle(open = true, id = 0) {
        Util.event.dispatch(Service.toggleEvent, { open, id });
    }
}

function MenuHeading({ title, triggerAdd }) {
    return (
        <div style={{ display: 'flex' }}>
            <div style={{ flex: 1 }}>
                <div>
                    <strong style={{ lineHeight: '24px' }}>{title}</strong>
                </div>
            </div>
            <div>
                <IconButton
                    icon={<PlusOutlined />}
                    type="primary"
                    title={title}
                    onClick={() => {
                        triggerAdd();
                    }}
                />
            </div>
        </div>
    );
}

export default function TaskField() {
    const setProjectId = useSetAtom(projectIdSt);
    const [id, setId] = useState(0);
    const [open, setOpen] = useState(false);
    const messages = getMessages();
    useEffect(() => {
        Util.event.listen(Service.toggleEvent, handleToggle);
        return () => {
            Util.event.remove(Service.toggleEvent, handleToggle);
        };
    }, []);

    const handleToggle = ({ detail: { open, id } }) => {
        if (!open) {
            setId(0);
            setOpen(false);
            return;
        }
        setProjectId(id);
        setId(id);
        setOpen(true);
    };

    return (
        <Drawer
            open={open}
            title={null}
            closeIcon={null}
            onClose={() => setOpen(false)}
            bodyStyle={{ padding: 10 }}
        >
            <MenuHeading
                title={messages.heading}
                triggerAdd={() => {
                    TaskFieldDialog.toggle(true);
                }}
            />
            <br />
            <TaskFieldTable projectId={id} />
        </Drawer>
    );
}

TaskField.displayName = 'TaskField';
TaskField.toggle = Service.toggle;
