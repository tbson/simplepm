import * as React from 'react';
import { t } from 'ttag';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Row, Col, Table, Button, Flex, Tooltip } from 'antd';
import { LockOutlined, UnlockOutlined } from '@ant-design/icons';
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

export default function UserTable() {
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
            .finally(() => {
                setInit(false);
            });
    };

    const handlePaging = (page) => {
        if (!page) {
            setPageParam({});
        } else {
            setPageParam({ page });
        }
    };

    const handleSearching = (keyword) => {
        setPageParam({});
        if (!keyword) {
            setSearchParam({});
        } else {
            setSearchParam({ q: keyword });
        }
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

    const onBulkDelete = (ids) => {
        const r = window.confirm(messages.deleteMultiple);
        if (!r) return;

        Util.toggleGlobalLoading(true);
        RequestUtil.apiCall(`${urls.crud}?ids=${ids.join(',')}`, {}, 'delete')
            .then(() => {
                setList([...list.filter((item) => !ids.includes(item.id))]);
            })
            .finally(() => Util.toggleGlobalLoading(false));
    };

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
            key: 'first_name',
            title: labels.first_name,
            dataIndex: 'first_name'
        },
        {
            key: 'last_name',
            title: labels.last_name,
            dataIndex: 'last_name'
        },
        {
            key: 'role_labels',
            title: labels.roles,
            dataIndex: 'role_labels',
            render: (text) => text.join(', ')
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
                        <EditBtn onClick={() => Dialog.toggle(true, record.id)} />
                    </PemCheck>
                    <PemCheck pem_group={PEM_GROUP} pem="delete">
                        <RemoveBtn onClick={() => onDelete(record.id)} />
                    </PemCheck>

                    <PemCheck pem_group={PEM_GROUP} pem="delete">
                        <Tooltip title={t`Lock`}>
                            <Button
                                danger={record.locked}
                                htmlType="button"
                                icon={record.locked ? <LockOutlined /> : <UnlockOutlined />}
                                size="small"
                                onClick={() => LockUserDialog.toggle(true, record.id)}
                            />
                        </Tooltip>
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
                        <RemoveSelectedBtn ids={ids} onClick={onBulkDelete} />
                    </PemCheck>
                </Col>
                <Col span={12} className="right"></Col>
            </Row>

            <SearchInput onChange={handleSearching} />

            <Table
                rowSelection={{
                    type: 'checkbox',
                    ...rowSelection
                }}
                onChange={handleTableChange}
                loading={init}
                columns={columns}
                dataSource={list}
                scroll={{ x: 1000 }}
                pagination={false}
            />
            <Pagination next={pages.next} prev={pages.prev} onChange={handlePaging} />
            <Dialog onChange={onChange} />
            <LockUserDialog onChange={onChange} />
        </div>
    );
}

UserTable.displayName = 'UserTable';
