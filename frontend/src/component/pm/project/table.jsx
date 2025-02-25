import * as React from 'react';
import { t } from 'ttag';
import { useEffect, useState } from 'react';
import { Link } from 'react-router';
import { useAtomValue } from 'jotai';
import { App, Row, Col, Table, Flex, Dropdown } from 'antd';
import {
    ProfileOutlined,
    SettingOutlined,
    DeleteOutlined,
    MoreOutlined
} from '@ant-design/icons';
import Pagination, { defaultPages } from 'component/common/table/pagination';
import SearchInput from 'component/common/table/search_input';
import {
    AddNewBtn
} from 'component/common/table/buttons';
import PemCheck from 'component/common/pem_check';
import Img from 'component/common/display/img';
import Util from 'service/helper/util';
import DictUtil from 'service/helper/dict_util';
import RequestUtil from 'service/helper/request_util';
import ProjectDialog from './dialog';
import TaskField from './task_field';
import { projectFilterSt } from 'component/pm/project/state';
import { urls, getLabels, getMessages, PEM_GROUP } from './config';

export default function ProjectTable() {
    const { notification } = App.useApp();
    const projectFilter = useAtomValue(projectFilterSt);
    const [searchParam, setSearchParam] = useState({});
    const [filterParam, setFilterParam] = useState({});
    const [sortParam, setSortParam] = useState({});
    const [pageParam, setPageParam] = useState({});
    const [init, setInit] = useState(false);
    const [list, setList] = useState([]);
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
                setList(Util.appendKeys(resp.data.items || []));
            })
            .catch(RequestUtil.displayError(notification))
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
            Util.event.dispatch('FETCH_BOOKMARK', {});
        } else {
            const index = list.findIndex((item) => item.id === id);
            data.key = data.id;
            list[index] = data;
            setList([...list]);
        }
    };

    const onDelete = (id) => {
        const r = window.confirm(messages.deleteOne);
        if (!r) {
            return;
        }

        Util.toggleGlobalLoading(true);
        RequestUtil.apiCall(`${urls.crud}${id}`, {}, 'delete')
            .then(() => {
                setList([...list.filter((item) => item.id !== id)]);
                Util.event.dispatch('FETCH_BOOKMARK', {});
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => Util.toggleGlobalLoading(false));
    };

    const getTableActions = (item) => {
        return {
            items: [
                {
                    key: 'setting',
                    label: t`Setting`,
                    icon: <SettingOutlined />,
                    onClick: () => {
                        ProjectDialog.toggle(true, item.id);
                    }
                },
                {
                    key: 'show_task_field',
                    label: t`Task field`,
                    icon: <ProfileOutlined />,
                    onClick: () => {
                        TaskField.toggle(true, item.id);
                    }
                },
                {
                    key: 'delete',
                    label: t`Delete`,
                    danger: true,
                    icon: <DeleteOutlined />,
                    onClick: () => {
                        onDelete(item.id);
                    }
                }
            ]
        };
    };

    const columns = [
        {
            key: 'avatar',
            title: labels.avatar,
            dataIndex: 'avatar',
            render: (value) => <Img src={value} width={30} height={30} />,
            width: 80
        },
        {
            key: 'title',
            title: labels.title,
            dataIndex: 'title',
            render: (value, record) => (
                <Link to={`/pm/project/${record.id}`}>{value}</Link>
            )
        },
        {
            key: 'status',
            title: labels.status,
            dataIndex: 'status',
            width: 120,
            filterMultiple: false,
            filters: projectFilter.status,
            onFilter: (value, record) => record.status === value
        },
        {
            key: 'action',
            title: '',
            fixed: 'right',
            width: 50,
            render: (_text, record) => (
                <Dropdown menu={getTableActions(record)} trigger={['click']}>
                    <MoreOutlined style={{ fontSize: '20px' }}/>
                </Dropdown>
            )
        }
    ];

    return (
        <div>
            <Row>
                <Col span={24} className="right">
                    <PemCheck pem_group={PEM_GROUP} pem="create">
                        <AddNewBtn onClick={() => ProjectDialog.toggle()} />
                    </PemCheck>
                </Col>
            </Row>

            <SearchInput onChange={handleSearching} />

            <Table
                onChange={handleTableChange}
                loading={init}
                columns={columns}
                dataSource={list}
                scroll={{ x: 1000 }}
                pagination={false}
            />
            <Pagination next={pages.next} prev={pages.prev} onChange={handlePaging} />
            <ProjectDialog onChange={onChange} />
            <TaskField />
        </div>
    );
}

ProjectTable.displayName = 'ProjectTable';
