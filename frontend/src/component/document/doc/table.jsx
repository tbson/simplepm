import * as React from 'react';
import { useState } from 'react';
import { createStyles } from 'antd-style';
import { useNavigate } from 'react-router-dom';
import { Button, Dropdown, List } from 'antd';
import {
    EditOutlined,
    DeleteOutlined,
    PlusOutlined,
    MoreOutlined,
    FileWordOutlined,
    UploadOutlined,
    LinkOutlined
} from '@ant-design/icons';
import NavUtil from 'service/helper/nav_util';
import { getStyles } from './style';

const testDocList = [
    {
        title: 'Document 1',
        url: 'https://www.google.com',
        type: 'DOC'
    },
    {
        title: 'Document 2',
        url: 'https://www.google.com',
        type: 'FILE'
    },
    {
        title: 'Document 3',
        url: 'https://www.google.com',
        type: 'LINK'
    }
];

export default function DocTable({ taskId, showControl = false }) {
    const navigate = useNavigate();
    const useStyle = getStyles(createStyles);
    const [documents, setDocuments] = useState(testDocList);
    const { styles } = useStyle();

    const navigateTo = NavUtil.navigateTo(navigate);

    const getDocumentMenuItems = () => {
        return {
            items: [
                {
                    key: 'document',
                    label: 'Doc',
                    icon: <FileWordOutlined />,
                    onClick: () => {
                        console.log('document');
                        navigateTo(`/pm/task/${taskId}/doc`);
                    }
                },
                {
                    key: 'file',
                    label: 'File',
                    icon: <UploadOutlined />,
                    onClick: () => {
                        console.log('document');
                    }
                },
                {
                    key: 'link',
                    label: 'Link',
                    icon: <LinkOutlined />,
                    onClick: () => {
                        console.log('document');
                    }
                }
            ]
        };
    };

    const getMessageMenuItems = (item) => {
        return {
            items: [
                {
                    key: 'edit',
                    label: 'Edit',
                    icon: <EditOutlined />,
                    onClick: () => {
                        console.log('edit', item);
                        setEditId(item.id);
                        setContent(item.content);
                    }
                },
                {
                    key: 'delete',
                    label: 'Delete',
                    danger: true,
                    icon: <DeleteOutlined />,
                    onClick: () => {
                        console.log('delete', item);
                        const r = window.confirm(
                            'Do you want to remove this document?'
                        );
                        if (!r) {
                            return;
                        }
                        deleteMessage(item.id);
                    }
                }
            ]
        };
    };

    const getDocumentIcon = (type) => {
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

    return (
        <div className={styles.document}>
            <div className={styles.chatHeading}>
                <div className="flex-item-remaining">
                    <div>
                        <strong>Documents</strong>
                    </div>
                </div>
                <div>
                    <Dropdown menu={getDocumentMenuItems()} trigger={['click']}>
                        <Button icon={<PlusOutlined />} />
                    </Dropdown>
                </div>
            </div>
            <List
                itemLayout="horizontal"
                size="small"
                dataSource={documents}
                renderItem={(item) => (
                    <List.Item>
                        <List.Item.Meta
                            avatar={getDocumentIcon(item.type)}
                            title={
                                <a href={item.url} target="_blank">
                                    {item.title}
                                </a>
                            }
                        />
                        {showControl ? (
                            <div>
                                <Dropdown
                                    menu={getMessageMenuItems(item)}
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
    );
}
