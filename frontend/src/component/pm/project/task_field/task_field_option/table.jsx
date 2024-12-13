import * as React from 'react';
import { useEffect } from 'react';
import { Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import useDraggableList from 'component/common/hook/use_draggable_list';
import TaskFieldOptionDialog from './dialog';

export default function TaskFieldOptionTable({ value, onChange }) {
    const [list, setList, DraggableListProvider, DraggableItem] = useDraggableList(
        value,
        (newItems) => handleSortEnd(newItems)
    );

    useEffect(() => {
        onChange(list);
    }, [list]);

    const handleSortEnd = (newItems) => {
        setList(newItems);
    };

    const handleChange = (data, id) => {
        if (id) {
            data.id = id;
            setList(list.map((record) => (record.id === id ? data : record)));
        } else {
            const tmpId = crypto.randomUUID();
            data.id = tmpId;
            setList([{ ...data }, ...list]);
        }
    };

    const handleDelete = (id) => {
        const newList = list.map((item) => {
            if (item.id === id) {
                item.deleted = true;
            }
            return item;
        });
        setList(newList);
    };

    return (
        <div>
            <Button
                type="dashed"
                onClick={() => TaskFieldOptionDialog.toggle(true)}
                block
                icon={<PlusOutlined />}
            >
                Add option
            </Button>
            <DraggableListProvider>
                {list
                    .filter((i) => !i.deleted)
                    .map((record) => (
                        <DraggableItem key={record.id} id={record.id}>
                            <div
                                className="pointer"
                                style={{ cursor: 'pointer' }}
                                onClick={() => {
                                    TaskFieldOptionDialog.toggle(true, record);
                                }}
                            >
                                <strong>{record.title}</strong>
                                <em
                                    style={{
                                        color: '#888',
                                        display: 'block',
                                        fontSize: '14px'
                                    }}
                                >
                                    {record.type}
                                </em>
                            </div>
                        </DraggableItem>
                    ))}
            </DraggableListProvider>
            <TaskFieldOptionDialog onChange={handleChange} onDelete={handleDelete} />
        </div>
    );
}

TaskFieldOptionTable.displayName = 'TaskFieldOptionTable';
