import React from 'react';
import { useState, useEffect } from 'react';
import { t } from 'ttag';
import { useAtom } from 'jotai';
import Util from 'service/helper/util';
import { Drawer, Divider } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { IconButton } from 'component/common/table/buttons';
import { projectOptionSt } from 'component/pm/project/state';
import TaskFieldTable from './table';
import TaskFieldDialog from './dialog';
import { TOGGLE_TASK_FIELD_EVENT } from './config';

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
            <div style={{ width: 40 }}>
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
    const [projectOption, setProjectOption] = useAtom(projectOptionSt);
    const [id, setId] = useState(0);
    const [open, setOpen] = useState(false);
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
        setProjectOption({ ...projectOption, project_id: id });
        setId(id);
        setOpen(true);
    };

    return (
        <Drawer
            open={open}
            title={null}
            closeIcon={null}
            onClose={() => setOpen(false)}
        >
            <MenuHeading
                title={t`Task fields`}
                triggerAdd={() => {
                    TaskFieldDialog.toggle(true);
                }}
            />
            <Divider />
            <TaskFieldTable projectId={id} />
        </Drawer>
    );
}

TaskField.displayName = 'TaskField';
TaskField.toggle = Service.toggle;
