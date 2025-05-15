import React, { useEffect, useState, useCallback } from 'react';
import { useAtomValue } from 'jotai';
import { App, Row, Col, Table } from 'antd';
import Pagination, { defaultPages } from 'component/common/table/pagination';
import SearchInput from 'component/common/table/search_input';
import {
    AddNewBtn,
    RemoveSelectedBtn,
    EditBtn,
    RemoveBtn
} from 'component/common/table/buttons';
import PemCheck from 'component/common/pem_check';
import Util from 'service/helper/util';
import DictUtil from 'service/helper/dict_util';
import RequestUtil from 'service/helper/request_util';
import Dialog from './dialog';
import { variableFilterSt } from 'component/config/variable/state';
import { urls, getLabels, getMessages, PEM_GROUP } from './config';

export default function VariableTable() {
    const { notification } = App.useApp();
    const variableFilter = useAtomValue(variableFilterSt);
    const [searchParam, setSearchParam] = useState({});
    const [filterParam, setFilterParam] = useState({});
    const [sortParam, setSortParam] = useState({});
    const [pageParam, setPageParam] = useState({});
    const [init, setInit] = useState(false);
    const [list, setList] = useState([]);
    const [ids, setIds] = useState([]);
    const [pages, setPages] = useState(defaultPages);
    const labels = getLabels();
    const messages = getMessages();

    useEffect(() => {
        getList();
    }, [searchParam, filterParam, sortParam, pageParam]);

    const getList = () => {
        setInit(true);
        const queryParam = {
            ...searchParam,
            ...filterParam,
            ...sortParam,
            ...pageParam
        };
        RequestUtil.apiCall(urls.crud, queryParam)
            .then((resp) => {
                setPages(resp.data.pages);
                const list = Util.appendKeys(resp.data.items);
                console.log(list);
                setList(list);
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                setInit(false);
            });
    };

    const handlePaging = useCallback((page) => {
        if (!page) {
            setPageParam({});
        } else {
            setPageParam({ page });
        }
    }, []);

    const handleSearch = useCallback((keyword) => {
        setPageParam({});
        if (!keyword) {
            setSearchParam({});
        } else {
            setSearchParam({ q: keyword });
        }
    }, []);

    const handleSortFilter = useCallback((_pagination, filters, sorter) => {
        const applyFilter = (filterObj) => {
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

        const applySort = (sortObj) => {
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

        setPageParam({});
        applyFilter(filters);
        applySort(sorter);
    }, []);

    const handleDataChange = useCallback(
        (data, id) => {
            if (!id) {
                setList([{ ...Util.appendKey(data) }, ...list]);
            } else {
                const index = list.findIndex((item) => item.id === id);
                list[index] = data;
                setList([...list]);
            }
        },
        [list]
    );

    const handleDelete = useCallback(
        (id) => {
            const r = window.confirm(messages.deleteOne);
            if (!r) return;

            Util.toggleGlobalLoading(true);
            RequestUtil.apiCall(`${urls.crud}${id}`, {}, 'delete')
                .then(() => {
                    setList([...list.filter((item) => item.id !== id)]);
                })
                .catch(RequestUtil.displayError(notification))
                .finally(() => Util.toggleGlobalLoading(false));
        },
        [list]
    );

    const handleBulkDelete = useCallback(() => {
        const r = window.confirm(messages.deleteMultiple);
        if (!r) return;

        Util.toggleGlobalLoading(true);
        RequestUtil.apiCall(`${urls.crud}?ids=${ids.join(',')}`, {}, 'delete')
            .then(() => {
                setList([...list.filter((item) => !ids.includes(item.id))]);
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => Util.toggleGlobalLoading(false));
    }, [ids, list]);

    const handleOpenAddEditDialog = useCallback((id) => {
        Dialog.toggle(true, id);
    }, []);

    const columns = [
        {
            key: 'key',
            title: labels.key,
            dataIndex: 'key',
            sorter: (a, b) => {
                return a.key.localeCompare(b.key);
            }
        },
        {
            key: 'value',
            title: labels.value,
            dataIndex: 'value'
        },
        {
            key: 'data_type',
            title: labels.data_type,
            dataIndex: 'data_type',
            width: 120,
            filterMultiple: false,
            filters: variableFilter.data_type,
            onFilter: (value, record) => record.data_type === value
        },
        {
            key: 'action',
            title: '',
            fixed: 'right',
            width: 90,
            render: (_text, record) => (
                <div className="flex-space">
                    <PemCheck pem_group={PEM_GROUP} pem="update">
                        <EditBtn value={record.id} onClick={handleOpenAddEditDialog} />
                    </PemCheck>
                    <PemCheck pem_group={PEM_GROUP} pem="delete">
                        <RemoveBtn value={record.id} onClick={handleDelete} />
                    </PemCheck>
                </div>
            )
        }
    ];

    const rowSelection = {
        type: 'checkbox',
        onChange: (ids) => {
            console.log(ids);
            setIds(ids);
        }
    };

    return (
        <div>
            <Row>
                <Col span={12}>
                    <PemCheck pem_group={PEM_GROUP} pem="deletelist">
                        <RemoveSelectedBtn value={ids} onClick={handleBulkDelete} />
                    </PemCheck>
                </Col>
                <Col span={12} className="right">
                    <PemCheck pem_group={PEM_GROUP} pem="create">
                        <AddNewBtn value={null} onClick={handleOpenAddEditDialog} />
                    </PemCheck>
                </Col>
            </Row>

            <SearchInput onChange={handleSearch} />

            <Table
                rowSelection={rowSelection}
                onChange={handleSortFilter}
                loading={init}
                columns={columns}
                dataSource={list}
                scroll={{ x: 1000 }}
                pagination={false}
            />
            <Pagination next={pages.next} prev={pages.prev} onChange={handlePaging} />
            <Dialog onChange={handleDataChange} />
        </div>
    );
}

VariableTable.displayName = 'VariableTable';
