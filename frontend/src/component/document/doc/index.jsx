import * as React from 'react';
import { useState } from 'react';
import { createStyles } from 'antd-style';
import { Button, Dropdown, List } from 'antd';
import {
    PlusOutlined,
    MoreOutlined,
    FileWordOutlined,
    UploadOutlined,
    LinkOutlined
} from '@ant-design/icons';
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

export default function Doc({ taskId }) {
    const useStyle = getStyles(createStyles);
    const [documents, setDocuments] = useState(testDocList);
    const { styles } = useStyle();

    const getDocumentMenuItems = () => {
        return {
            items: [
                {
                    key: 'document',
                    label: 'Documnet',
                    icon: <FileWordOutlined />,
                    onClick: () => {
                        console.log('document');
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
                        <div>
                            <MoreOutlined style={{ fontSize: '20px' }} />
                        </div>
                    </List.Item>
                )}
            />
        </div>
    );
}
