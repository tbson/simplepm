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
    const [statusList, setStatusList] = useState([]);
    const taskOption = useAtomValue(taskOptionSt);
    const [filterParam, setFilterParam] = useState({});
    const [sortParam, setSortParam] = useState({});
    const [pageParam, setPageParam] = useState({});
    const [init, setInit] = useState(false);
    const [list, setList] = useState([]);
    const labels = getLabels();
    const messages = getMessages();
    const statusMap = taskOption.status.reduce((acc, item) => {
        acc[item.value] = item.label;
        return acc;
    }, {});

    useEffect(() => {
        if (taskOption.loaded) {
            getList();
        }
    }, [taskOption.loaded, filterParam, sortParam, pageParam]);

    const getList = () => {
        setInit(true);
        const queryParam = {
            ...filterParam,
            ...sortParam,
            ...pageParam
        };
        RequestUtil.apiCall(urls.crud, { ...queryParam, project_id })
            .then((resp) => {
                const list = resp.data.map((item) => {
                    return {
                        id: item.id,
                        title: item.title,
                        status: item.status.id
                    };
                });
                setList(list);
                const statusList = taskOption.status.map((status, index) => {
                    return {
                        id: status.value,
                        title: status.label,
                        order: index + 1
                    };
                });
                setStatusList(statusList);
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

    const handleAdd = (status) => {
        Dialog.toggle();
    };

    const handleChange = (result) => {
        result.project_id = project_id;
        console.log(result);
        RequestUtil.apiCall(urls.reorder, result, 'put')
            .then((resp) => {
                console.log(resp);
            })
            .catch((err) => {
                console.log(err);
            });
    };

    return (
        <div>
            <Kanban tasks={list} statusList={statusList} onChange={handleChange} />
            <Dialog onChange={onChange} />
        </div>
    );
}

TaskKanban.displayName = 'TaskKanban';
