import * as React from 'react';
import { useEffect, useState } from 'react';
import { Skeleton } from 'antd';
import RequestUtil from 'service/helper/request_util';
import Util from 'service/helper/util';
import useDraggableList from 'component/common/hook/use_draggable_list';
import Dialog from './dialog';
import { urls, getLabels } from './config';

export default function TaskfieldTable({ projectId }) {
    const [init, setInit] = useState(false);
    const [list, setList, DraggableListProvider, DraggableItem] = useDraggableList(
        [],
        (newItems) => handleSortEnd(newItems)
    );

    useEffect(() => {
        getList();
    }, [projectId]);

    const handleSortEnd = (newItems) => {
        console.log(newItems);
        const items = newItems.map((id, index) => {
            return { id, order: index + 1 };
        });
        const payload = { items, project_id: projectId };

        RequestUtil.apiCall(urls.reorder, payload, 'put')
            .then(() => {
                console.log('reorder success');
            })
            .catch((err) => {
                console.log(err);
            });
    };

    const getList = () => {
        setInit(true);
        const params = { project_id: projectId };
        RequestUtil.apiCall(urls.crud, params)
            .then((resp) => {
                setList(Util.appendKeys(resp.data));
            })
            .finally(() => {
                setInit(false);
            });
    };

    const onChange = (data, id) => {
        if (!id) {
            setList([{ ...Util.appendKey(data) }, ...list]);
        } else {
            const index = list.findIndex((item) => item.id === id);
            data.key = data.id;
            list[index] = data;
            setList([...list]);
        }
    };

    const onDelete = (id) => {
        setList([...list.filter((item) => item.id !== id)]);
    };

    return (
        <>
            <Skeleton loading={init} active />
            <DraggableListProvider>
                {list.map((record) => (
                    <DraggableItem key={record.id} id={record.id}>
                        <div
                            className="pointer"
                            style={{ cursor: 'pointer' }}
                            onClick={() => Dialog.toggle(true, record.id)}
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

            <Dialog onChange={onChange} onDelete={onDelete} />
        </>
    );
}
