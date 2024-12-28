import * as React from 'react';
import { useEffect, useState } from 'react';
import { useAtomValue } from 'jotai';
import { Table } from 'antd';
import { EditBtn, RemoveBtn } from 'component/common/table/buttons';
import PemCheck from 'component/common/pem_check';
import Kanban from 'component/common/kanban';
import Util from 'service/helper/util';
import DictUtil from 'service/helper/dict_util';
import RequestUtil from 'service/helper/request_util';
import Dialog from './dialog';
import { taskOptionSt } from 'component/pm/task/state';
import { urls, getLabels, getMessages, PEM_GROUP } from './config';

export default function TaskKanban({ project_id }) {
    const taskOption = useAtomValue(taskOptionSt);
    const [filterParam, setFilterParam] = useState({});
    const [sortParam, setSortParam] = useState({});
    const [pageParam, setPageParam] = useState({});
    const [init, setInit] = useState(false);
    const [list, setList] = useState([]);
    const labels = getLabels();
    const messages = getMessages();

    useEffect(() => {
        getList();
    }, [filterParam, sortParam, pageParam]);

    const getList = () => {
        setInit(true);
        const queryParam = {
            ...filterParam,
            ...sortParam,
            ...pageParam
        };
        RequestUtil.apiCall(urls.crud, { ...queryParam, project_id })
            .then((resp) => {
                setList(Util.appendKeys(resp.data));
            })
            .finally(() => {
                setInit(false);
            });
    };

    const handleFiltering = (filterObj) => {
        if (DictUtil.isEmpty(filterObj)) {
            setFilterParam({});
        } else {
            setFilterParam(
                Object.entries(filterObj).reduce((acc, [key, value]) => {
                    if (!value || value.length === 0) {
                        return acc;
                    }
                    acc[key] = value[0];
                    return acc;
                }, {})
            );
        }
    };

    const handleSorting = (sortObj) => {
        if (DictUtil.isEmpty(sortObj)) {
            return setSortParam({});
        }
        if (!sortObj.field) {
            return setSortParam({});
        }
        const sign = sortObj.order === 'descend' ? '-' : '';
        setSortParam({
            order: `${sign}${sortObj.field}`
        });
    };

    const handleTableChange = (_pagination, filters, sorter) => {
        setPageParam({});
        handleFiltering(filters);
        handleSorting(sorter);
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
        const r = window.confirm(messages.deleteOne);
        if (!r) return;

        Util.toggleGlobalLoading(true);
        RequestUtil.apiCall(`${urls.crud}${id}`, {}, 'delete')
            .then(() => {
                setList([...list.filter((item) => item.id !== id)]);
            })
            .finally(() => Util.toggleGlobalLoading(false));
    };

    const columns = [
        {
            key: 'title',
            title: labels.title,
            dataIndex: 'title'
        },
        {
            key: 'feature_id',
            title: labels.feature,
            dataIndex: 'feature_id',
            width: 120
        },
        {
            key: 'action',
            title: '',
            fixed: 'right',
            width: 90,
            render: (_text, record) => (
                <div className="flex-space">
                    <PemCheck pem_group={PEM_GROUP} pem="update">
                        <EditBtn onClick={() => Dialog.toggle(true, record.id)} />
                    </PemCheck>
                    <PemCheck pem_group={PEM_GROUP} pem="delete">
                        <RemoveBtn onClick={() => onDelete(record.id)} />
                    </PemCheck>
                </div>
            )
        }
    ];
    if (!taskOption.loaded) {
        return null;
    }
    return (
        <div>
            <Kanban statusList={taskOption.status.map((i) => i.label)} data={list} />
            <Dialog onChange={onChange} />
        </div>
    );
}

TaskKanban.displayName = 'TaskKanban';
