import * as React from 'react';
import { useEffect } from 'react';
import { Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import DateUtil from 'service/helper/date_util';
import useDraggableList from 'component/common/hook/use_draggable_list';
import TaskFieldOptionDialog from './dialog';

const ITEM_STATUS = {
    EXISTING: 'EXISTING',
    CREATED: 'CREATED',
    UPDATED: 'UPDATED',
    DELETED: 'DELETED'
};

export default function TaskFieldOptionTable({ value, onChange }) {
    const [list, setList, DraggableListProvider, DraggableItem] = useDraggableList(
        value,
        (newItems) => handleSortEnd(newItems)
    );

    useEffect(() => {
        onChange(
            list.map((item) => {
                if (!item.fe_status) {
                    item.fe_status = ITEM_STATUS.EXISTING;
                }
                return item;
            })
        );
    }, [list]);

    const handleSortEnd = (sortedItems) => {
        const newItems = sortedItems.map((item, index) => {
            item.order = index + 1;
            if (item.fe_status === ITEM_STATUS.EXISTING) {
                item.fe_status = ITEM_STATUS.UPDATED;
            }
            return item;
        });
        setList(newItems);
    };

    const handleChange = (data, id) => {
        if (id) {
            data.id = id;
            if (data.fe_status === ITEM_STATUS.EXISTING) {
                data.fe_status = ITEM_STATUS.UPDATED;
            }
            setList(list.map((record) => (record.id === id ? data : record)));
        } else {
            const tmpId = DateUtil.currentTimestamp();
            data.id = tmpId;
            data.order = list.length + 1;
            data.fe_status = ITEM_STATUS.CREATED;
            setList([...list, { ...data }]);
        }
    };

    const handleDelete = (id) => {
        const item = list.find((i) => i.id === id);
        if (item.fe_status === ITEM_STATUS.CREATED) {
            setList(list.filter((i) => i.id !== id));
            return;
        }
        const newList = list.map((item) => {
            if (item.id === id) {
                item.fe_status = ITEM_STATUS.DELETED;
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
                    .filter((i) => i.fe_status !== ITEM_STATUS.DELETED)
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
