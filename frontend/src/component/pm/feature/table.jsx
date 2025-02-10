import * as React from 'react';
import { useEffect, useState } from 'react';
import { useSetAtom } from 'jotai';
import { Link } from 'react-router';
import { App, Button, Badge } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import useDraggableList from 'component/common/hook/use_draggable_list';
import { featureColorSt } from './state';
import Dialog from './dialog';
import { urls, getMessages } from './config';

export default function FeatureTable({ projectId }) {
    const { notification } = App.useApp();
    const setFeatureColor = useSetAtom(featureColorSt);
    const [init, setInit] = useState(false);
    const [list, setList, DraggableListProvider, DraggableItem] = useDraggableList(
        [],
        (newItems) => handleSortEnd(newItems)
    );
    const messages = getMessages();

    useEffect(() => {
        getList();
    }, [projectId]);

    const getList = () => {
        setInit(true);
        RequestUtil.apiCall(urls.crud, { project_id: projectId })
            .then((resp) => {
                setList(Util.appendKeys(resp.data));
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                setInit(false);
            });
    };

    const handleChange = (data, id) => {
        if (!id) {
            setList([{ ...Util.appendKey(data) }, ...list]);
        } else {
            const index = list.findIndex((item) => item.id === id);
            const item = list[index];
            if (item.color !== data.color) {
                setFeatureColor({ featureId: data.id, color: data.color });
            }
            data.key = data.id;
            list[index] = data;
            setList([...list]);
        }
    };

    const handleDelete = (id) => {
        const r = window.confirm(messages.deleteOne);
        if (!r) return;

        Util.toggleGlobalLoading(true);
        RequestUtil.apiCall(`${urls.crud}${id}`, {}, 'delete')
            .then(() => {
                setList([...list.filter((item) => item.id !== id)]);
                Dialog.toggle(false);
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                Util.toggleGlobalLoading(false);
            });
    };

    const handleSortEnd = (newItems) => {
        const items = newItems.map((item, index) => {
            return { id: item.id, order: index + 1 };
        });
        const payload = { items, project_id: projectId };

        RequestUtil.apiCall(urls.reorder, payload, 'put')
            .then(() => {
                console.log('reorder success');
            })
            .catch(RequestUtil.displayError(notification));
    };

    return (
        <div>
            <DraggableListProvider
                layout="horizontal"
                fixedComponent={
                    <Button
                        className="card"
                        onClick={() => Dialog.toggle()}
                        icon={<PlusOutlined />}
                        size="large"
                    >
                        Feature
                    </Button>
                }
            >
                {list.map((record) => (
                    <DraggableItem key={record.id} id={record.id} color={record.color}>
                        <Badge count={0} offset={[10, -10]}>
                            <Link to={`/pm/task/message/${projectId}/${record.id}`}>
                                <div
                                    className="pointer"
                                    onClick={() => Dialog.toggle(true, record.id)}
                                >
                                    {record.title}
                                </div>
                            </Link>
                        </Badge>
                    </DraggableItem>
                ))}
            </DraggableListProvider>
            <Dialog onChange={handleChange} onDelete={handleDelete} />
        </div>
    );
}

FeatureTable.displayName = 'FeatureTable';
