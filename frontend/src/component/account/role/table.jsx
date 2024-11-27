import * as React from 'react';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Row, Col, Table, Flex } from 'antd';
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
import FormUtil from 'service/helper/form_util';
import RequestUtil from 'service/helper/request_util';
import Dialog from './dialog';
import { urls, getLabels, getMessages, PEM_GROUP } from './config';

export default function RoleTable() {
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
            }).catch((err) => {
                FormUtil.setFormErrors()(err.response.data);
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
            key: 'title',
            title: labels.title,
            dataIndex: 'title',
            sorter: (a, b) => {
                return a.title.localeCompare(b.title);
            }
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
                <Col span={12} className="right">
                    <PemCheck pem_group={PEM_GROUP} pem="create">
                        <AddNewBtn onClick={() => Dialog.toggle()} />
                    </PemCheck>
                </Col>
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
        </div>
    );
}

RoleTable.displayName = 'RoleTable';
