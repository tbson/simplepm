import * as React from 'react';
import { useState } from 'react';
import { createStyles } from 'antd-style';
import { useNavigate } from 'react-router';
import { App, Button, Dropdown, List } from 'antd';
import {
    DeleteOutlined,
    PlusOutlined,
    MoreOutlined,
    FileWordOutlined,
    UploadOutlined,
    LinkOutlined
} from '@ant-design/icons';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import FormUtil from 'service/helper/form_util';
import NavUtil from 'service/helper/nav_util';
import DocLinkDialog from './link_dialog';
import { getStyles } from './style';
import { urls } from './config';

export default function DocTable({ taskId, showControl = false }) {
    const { notification } = App.useApp();
    const navigate = useNavigate();
    const useStyle = getStyles(createStyles);
    const [list, setList] = useState([]);
    const { styles } = useStyle();

    React.useEffect(() => {
        getList();
    }, [taskId]);

    const getList = () => {
        const queryParam = {
            task_id: taskId
        };
        RequestUtil.apiCall(urls.crud, queryParam)
            .then((resp) => {
                // setPages(resp.data.pages);
                setList(Util.appendKeys(resp.data.items));
            })
            .catch(RequestUtil.displayError(notification));
    };

    const navigateTo = NavUtil.navigateTo(navigate);

    const getDocCreateOptions = () => {
        return {
            items: [
                {
                    key: 'document',
                    label: 'Doc',
                    icon: <FileWordOutlined />,
                    onClick: () => {
                        navigateTo(`/pm/task/${taskId}/doc`);
                    }
                },
                {
                    key: 'file',
                    label: 'File',
                    icon: <UploadOutlined />,
                    onClick: () => {
                        FormUtil.getFile().then((file) => {
                            const fileName = file.name;
                            const fileSize = file.size;
                            const fileType = file.type;
                            const payload = {
                                type: 'FILE',
                                task_id: taskId,
                                title: fileName,
                                file_name: fileName,
                                file_size: fileSize,
                                file_type: fileType,
                                file_url: file
                            };
                            RequestUtil.apiCall(urls.crud, payload, 'post')
                                .then(() => {
                                    getList();
                                })
                                .catch(RequestUtil.displayError(notification));
                        });
                    }
                },
                {
                    key: 'link',
                    label: 'Link',
                    icon: <LinkOutlined />,
                    onClick: () => {
                        DocLinkDialog.toggle(true);
                    }
                }
            ]
        };
    };

    const getDocIcon = (type) => {
        if (type === 'DOC') {
            return <FileWordOutlined />;
        }
        if (type === 'FILE') {
            return <UploadOutlined />;
        }
        if (type === 'LINK') {
            return <LinkOutlined />;
        }
        return null;
    };

    const getTableActions = (item) => {
        return {
            items: [
                {
                    key: 'delete',
                    label: 'Delete',
                    danger: true,
                    icon: <DeleteOutlined />,
                    onClick: () => {
                        const r = window.confirm(
                            'Do you want to remove this document?'
                        );
                        if (!r) {
                            return;
                        }
                        RequestUtil.apiCall(`${urls.crud}${item.id}`, {}, 'delete')
                            .then(() => {
                                getList();
                            })
                            .catch(RequestUtil.displayError(notification));
                    }
                }
            ]
        };
    };

    return (
        <>
            <div className={styles.document}>
                <div className={styles.chatHeading}>
                    <div className="flex-item-remaining">
                        <div>
                            <strong>Documents</strong>
                        </div>
                    </div>
                    <div>
                        <Dropdown menu={getDocCreateOptions()} trigger={['click']}>
                            <Button icon={<PlusOutlined />} />
                        </Dropdown>
                    </div>
                </div>
                <List
                    itemLayout="horizontal"
                    size="small"
                    dataSource={list}
                    renderItem={(item) => (
                        <List.Item>
                            <List.Item.Meta
                                avatar={getDocIcon(item.type)}
                                title={
                                    <a href={item.url} target="_blank">
                                        {item.title}
                                    </a>
                                }
                                onClick={() => {
                                    if (item.type === 'FILE') {
                                        return window.open(item.file_url);
                                    }
                                    if (item.type === 'LINK') {
                                        return window.open(item.link);
                                    }
                                    navigateTo(`/pm/task/${taskId}/doc/${item.id}`);
                                }}
                            />
                            {showControl ? (
                                <div>
                                    <Dropdown
                                        menu={getTableActions(item)}
                                        trigger={['click']}
                                    >
                                        <MoreOutlined style={{ fontSize: '20px' }} />
                                    </Dropdown>
                                </div>
                            ) : null}
                        </List.Item>
                    )}
                />
            </div>
            <DocLinkDialog onChange={getList} />
        </>
    );
}
