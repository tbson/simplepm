import * as React from 'react';
import { Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import useDraggableList from 'component/common/hook/use_draggable_list';
import Dialog from './dialog';

export default function TaskFieldOptionTable({ initList }) {
    const [list, setList, DraggableListProvider, DraggableItem] = useDraggableList(
        initList,
        (newItems) => handleSortEnd(newItems)
    );

    const handleSortEnd = (newItems) => {
        setList(newItems);
    };

    const handleChange = (data, id) => {
        if (id) {
            setList(list.map((record) => (record.id === id ? data : record)));
        } else {
            const tmpId = crypto.randomUUID();
            data.id = tmpId;
            setList([{ ...data }, ...list]);
        }
    };

    return (
        <div>
            <Button
                type="dashed"
                onClick={() => Dialog.toggle(true)}
                block
                icon={<PlusOutlined />}
            >
                Add option
            </Button>
            <DraggableListProvider>
                {list.map((record) => (
                    <DraggableItem key={record.id} id={record.id}>
                        <div
                            className="pointer"
                            style={{ cursor: 'pointer' }}
                            onClick={() => Dialog.toggle(true, record)}
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
            <Dialog onChange={handleChange} />
        </div>
    );
}

TaskFieldOptionTable.displayName = 'TaskFieldOptionTable';
