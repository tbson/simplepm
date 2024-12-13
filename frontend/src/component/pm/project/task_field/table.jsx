import * as React from 'react';
import { useEffect, useState } from 'react';
import { Skeleton } from 'antd';
import RequestUtil from 'service/helper/request_util';
import Util from 'service/helper/util';
import useDraggableList from 'component/common/hook/use_draggable_list';
import Dialog from './dialog';
import { urls, TOGGLE_DIALOG_EVENT } from './config';

export class Service {
    static get toggleEvent() {
        return TOGGLE_DIALOG_EVENT;
    }

    static toggle(open = true, id = 0) {
        Util.event.dispatch(Service.toggleEvent, { open, id });
    }
}

export default function TaskFieldTable({ projectId }) {
    const [data, setData] = useState({});
    const [open, setOpen] = useState(false);
    const [id, setId] = useState(0);

    const [init, setInit] = useState(false);
    const [list, setList, DraggableListProvider, DraggableItem] = useDraggableList(
        [],
        (newItems) => handleSortEnd(newItems)
    );

    const handleToggle = ({ detail: { open, id } }) => {
        if (!open) {
            return setOpen(false);
        }
        setId(id);
        if (id) {
            Util.toggleGlobalLoading();
            RequestUtil.apiCall(`${urls.crud}${id}`)
                .then((resp) => {
                    setData(resp.data);
                    setOpen(true);
                })
                .finally(() => Util.toggleGlobalLoading(false));
        } else {
            setData({});
            setOpen(true);
        }
    };

    useEffect(() => {
        Util.event.listen(Service.toggleEvent, handleToggle);
        return () => {
            Util.event.remove(Service.toggleEvent, handleToggle);
        };
    }, []);

    useEffect(() => {
        getList();
    }, [projectId]);

    const handleSortEnd = (newItems) => {
        const items = newItems.map((item, index) => {
            return { id: item.id, order: index + 1 };
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

    const renderTable = () => {
        return (
            <>
                <Skeleton loading={init} active />
                <DraggableListProvider>
                    {list.map((record) => (
                        <DraggableItem key={record.id} id={record.id}>
                            <div
                                className="pointer"
                                style={{ cursor: 'pointer' }}
                                onClick={() => Service.toggle(true, record.id)}
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
            </>
        );
    };

    const renderDialog = () => {
        return (
            <Dialog
                id={id}
                data={data}
                toggle={Service.toggle}
                onChange={onChange}
                onDelete={onDelete}
            />
        );
    };

    return open ? renderDialog() : renderTable();
}

TaskFieldTable.displayName = 'TaskFieldTable';
TaskFieldTable.toggle = Service.toggle;
