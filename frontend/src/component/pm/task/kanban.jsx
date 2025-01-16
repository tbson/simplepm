import * as React from 'react';
import { useEffect, useState } from 'react';
import { App } from 'antd';
import { useAtomValue } from 'jotai';
import { useNavigate } from 'react-router-dom';
import PemCheck from 'component/common/pem_check';
import Kanban, { REORDER_TASK } from 'component/common/kanban';
import Util from 'service/helper/util';
import NavUtil from 'service/helper/nav_util';
import DictUtil from 'service/helper/dict_util';
import RequestUtil from 'service/helper/request_util';
import Dialog from './dialog';
import { taskOptionSt } from 'component/pm/task/state';
import { featureColorSt } from 'component/pm/feature/state';
import { urls, getLabels, getMessages, PEM_GROUP } from './config';

export default function TaskKanban({ projectId }) {
    const { notification } = App.useApp();
    const navigate = useNavigate();
    const taskOption = useAtomValue(taskOptionSt);
    const featureColor = useAtomValue(featureColorSt);
    const [statusList, setStatusList] = useState([]);
    const [filterParam, setFilterParam] = useState({});
    const [sortParam, setSortParam] = useState({});
    const [init, setInit] = useState(false);
    const [list, setList] = useState([]);
    const labels = getLabels();
    const messages = getMessages();

    const navigateTo = NavUtil.navigateTo(navigate);

    useEffect(() => {
        if (taskOption.loaded) {
            getList();
        }
    }, [taskOption.loaded, filterParam, sortParam]);

    useEffect(() => {
        if (!featureColor.featureId) return;
        const newList = list.map((item) => {
            /*
            if (item.featureId === featureColor.featureId) {
                item.color = featureColor.color;
            }
            */
            return item;
        });
        setList(newList);
    }, [featureColor]);

    const getList = () => {
        setInit(true);
        const queryParam = {
            ...filterParam,
            ...sortParam
        };
        RequestUtil.apiCall(urls.crud, { ...queryParam, project_id: projectId })
            .then((resp) => {
                const list = resp.data.map((item) => {
                    return {
                        id: item.id,
                        title: item.title,
                        status: item.status.id,
                        // featureId: item.feature.id,
                        // color: item.feature.color,
                        order: item.order
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
            .catch(RequestUtil.displayError(notification))
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

    const handleChange = (data, id) => {
        data.status = data.status.id;
        data.color = data.feature.color;
        if (!id) {
            const newList = [{ ...Util.appendKey(data) }, ...list];
            newList.sort((a, b) => a.order - b.order);
            setList(newList);
        } else {
            const index = list.findIndex((item) => item.id === id);
            data.key = data.id;
            list[index] = data;
            setList([...list]);
        }
    };

    const handleDelete = (id) => {
        const r = window.confirm(messages.deleteOne);
        if (!r) {
            return;
        }

        Util.toggleGlobalLoading(true);
        RequestUtil.apiCall(`${urls.crud}${id}`, {}, 'delete')
            .then(() => {
                setList([...list.filter((item) => item.id !== id)]);
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                Dialog.toggle(false);
                Util.toggleGlobalLoading(false);
            });
    };

    const handleAdd = (status) => {
        Dialog.toggle(true, 0, status);
    };

    const handleView = (id) => {
        // Dialog.toggle(true, id);
        navigateTo(`/pm/task/${projectId}/${id}`);
    };

    const handleReorder = (type, data) => {
        data.project_id = projectId;
        const endpoint = type === REORDER_TASK ? urls.reorder : urls.reorderStatus;
        return RequestUtil.apiCall(endpoint, data, 'put').catch(
            RequestUtil.displayError(notification)
        );
    };

    return (
        <div>
            <Kanban
                tasks={list}
                statusList={statusList}
                onReorder={handleReorder}
                onAdd={handleAdd}
                onView={handleView}
            />
            <Dialog onChange={handleChange} onDelete={handleDelete} />
        </div>
    );
}

TaskKanban.displayName = 'TaskKanban';
