import * as React from 'react';
import { useEffect, useState } from 'react';
import { Button, Badge } from 'antd';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import useDraggableList from 'component/common/hook/use_draggable_list';
import Dialog from './dialog';
import { urls } from './config';

export default function FeatureTable({ project_id }) {
    const [init, setInit] = useState(false);
    const [list, setList, DraggableListProvider, DraggableItem] = useDraggableList(
        [],
        (newItems) => handleSortEnd(newItems)
    );

    useEffect(() => {
        getList();
    }, [project_id]);

    const getList = () => {
        setInit(true);
        RequestUtil.apiCall(urls.crud, { project_id })
            .then((resp) => {
                setList(Util.appendKeys(resp.data));
            })
            .finally(() => {
                setInit(false);
            });
    };

    const handleChange = (data, id) => {
        if (!id) {
            setList([{ ...Util.appendKey(data) }, ...list]);
        } else {
            const index = list.findIndex((item) => item.id === id);
            data.key = data.id;
            list[index] = data;
            setList([...list]);
        }
    };

    const handleSortEnd = (newItems) => {
        const items = newItems.map((item, index) => {
            return { id: item.id, order: index + 1 };
        });
        const payload = { items, project_id };

        RequestUtil.apiCall(urls.reorder, payload, 'put')
            .then(() => {
                console.log('reorder success');
            })
            .catch((err) => {
                console.log(err);
            });
    };

    return (
        <div>
            <div>
                <Button onClick={() => Dialog.toggle()}>Add</Button>
            </div>
            <DraggableListProvider layout="horizontal">
                {list.map((record) => (
                    <DraggableItem key={record.id} id={record.id}>
                        <Badge count={5} offset={[10, -10]}>
                            <div
                                className="pointer"
                                style={{ cursor: 'pointer' }}
                                onClick={() => Dialog.toggle(true, record.id)}
                            >
                                <strong>{record.title}</strong>
                            </div>
                        </Badge>
                    </DraggableItem>
                ))}
            </DraggableListProvider>
            <Dialog onChange={handleChange} />
        </div>
    );
}

FeatureTable.displayName = 'FeatureTable';
