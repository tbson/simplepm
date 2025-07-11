import React, { useEffect, useState, useCallback } from 'react';
import { t } from 'ttag';
import { useParams } from 'react-router';
import { App, Row, Col, Table, Button, Flex, Tooltip } from 'antd';
import { LockOutlined, UnlockOutlined, GithubOutlined } from '@ant-design/icons';
import Pagination, { defaultPages } from 'component/common/table/pagination';
import SearchInput from 'component/common/table/search_input';
import { RemoveSelectedBtn, EditBtn, RemoveBtn } from 'component/common/table/buttons';
import PemCheck from 'component/common/pem_check';
import Util from 'service/helper/util';
import DictUtil from 'service/helper/dict_util';
import RequestUtil from 'service/helper/request_util';
import Dialog from './dialog';
import LockUserDialog from './lock_user';
import { urls, getLabels, getMessages, PEM_GROUP } from './config';

function LockButton({ locked, value, onClick }) {
    const handleClick = useCallback(() => {
        onClick(value);
    }, [value]);

    return (
        <Tooltip title={t`Lock`}>
            <Button
                danger={locked}
                htmlType="button"
                icon={locked ? <LockOutlined /> : <UnlockOutlined />}
                size="small"
                onClick={handleClick}
            />
        </Tooltip>
    );
}

export default function UserTable() {
    const { notification } = App.useApp();
    const { tenant_id } = useParams();
    const defaultFilterParam = tenant_id ? { tenant_id } : {};
    const [searchParam, setSearchParam] = useState({});
    const [filterParam, setFilterParam] = useState(defaultFilterParam);
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
                setList(Util.appendKeys(resp.data.items));
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
                data.key = data.id;
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

    const handleToggleLock = useCallback((id) => {
        LockUserDialog.toggle(true, id);
    }, []);

    const columns = [
        {
            key: 'email',
            title: labels.email,
            dataIndex: 'email',
            sorter: (a, b) => {
                return a.key.localeCompare(b.key);
            }
        },
        {
            key: 'mobile',
            title: labels.mobile,
            dataIndex: 'mobile'
        },
        {
            key: 'full_name',
            title: labels.full_name,
            dataIndex: 'full_name'
        },
        {
            key: 'git_account',
            title: labels.git_account,
            render: (_text, record) => {
                return (
                    <div>
                        {record.github_username ? (
                            <div>
                                <GithubOutlined /> &nbsp;
                                <a
                                    href={`https://github.com/${record.github_username}`}
                                    target="_blank"
                                    rel="noopener noreferrer"
                                >
                                    {record.github_username}
                                </a>
                            </div>
                        ) : null}
                    </div>
                );
            }
        },
        {
            key: 'role_labels',
            title: labels.roles,
            dataIndex: 'role_labels',
            render: (text) => text && text.join(', ')
        },
        {
            key: 'admin',
            title: labels.admin,
            dataIndex: 'admin',
            render: (text) => (text ? 'Yes' : 'No'),
            width: 90
        },
        {
            key: 'action',
            title: '',
            fixed: 'right',
            width: 90,
            render: (_text, record) => (
                <Flex wrap gap={5} justify="flex-end">
                    <PemCheck pem_group={PEM_GROUP} pem="update">
                        <EditBtn value={record.id} onClick={handleOpenAddEditDialog} />
                    </PemCheck>
                    <PemCheck pem_group={PEM_GROUP} pem="delete">
                        <RemoveBtn value={record.id} onClick={handleDelete} />
                    </PemCheck>

                    <PemCheck pem_group={PEM_GROUP} pem="delete">
                        <LockButton
                            locked={record.locked}
                            value={record.id}
                            onClick={handleToggleLock}
                        />
                    </PemCheck>
                </Flex>
            )
        }
    ];

    const rowSelection = {
        onChange: (ids) => {
            setIds(ids);
        }
    };

    return (
        <div>
            <Row>
                <Col span={12}>
                    <PemCheck pem_group={PEM_GROUP} pem="delete_list">
                        <RemoveSelectedBtn value={ids} onClick={handleBulkDelete} />
                    </PemCheck>
                </Col>
            </Row>

            <SearchInput onChange={handleSearch} />

            <Table
                rowSelection={{
                    type: 'checkbox',
                    ...rowSelection
                }}
                onChange={handleSortFilter}
                loading={init}
                columns={columns}
                dataSource={list}
                scroll={{ x: 1000 }}
                pagination={false}
            />
            <Pagination next={pages.next} prev={pages.prev} onChange={handlePaging} />
            <Dialog onChange={handleDataChange} />
            <LockUserDialog onChange={handleDataChange} />
        </div>
    );
}

UserTable.displayName = 'UserTable';
