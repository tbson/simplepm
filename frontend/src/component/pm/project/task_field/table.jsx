import * as React from 'react';
import { useEffect, useState } from 'react';
import { Table } from 'antd';
import RequestUtil from 'service/helper/request_util';
import Util from 'service/helper/util';
import Dialog from './dialog';
import { urls, getLabels } from './config';

export default function TaskfieldTable({ projectId }) {
    const [init, setInit] = useState(false);
    const [list, setList] = useState([]);
    const labels = getLabels();

    useEffect(() => {
        getList();
    }, [projectId]);

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
    const columns = [
        {
            key: 'title',
            title: labels.title,
            dataIndex: 'title',
            render: (text, record) => (
                <div className="pointer" onClick={() => Dialog.toggle(true, record.id)}>
                    <div>{text}</div>
                    <div><em>{record.type}</em></div>
                </div>
            )
        }
    ];

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
    }

    return (
        <>
            <Table
                loading={init}
                columns={columns}
                dataSource={list}
                pagination={false}
            />
            <Dialog onChange={onChange} onDelete={onDelete} />
        </>
    );
}
