import * as React from 'react';
import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { App, Badge, Button, Flex, Avatar } from 'antd';
import { Attachments, Bubble, Conversations, Sender } from '@ant-design/x';
import { createStyles } from 'antd-style';
import {
    CloudUploadOutlined,
    PaperClipOutlined,
    EditOutlined
} from '@ant-design/icons';
import Util from 'service/helper/util';
import NavUtil from 'service/helper/nav_util';
import RequestUtil from 'service/helper/request_util';
import SocketUtil from 'service/helper/socket_util';
import StorageUtil from 'service/helper/storage_util';
import TaskDialog from 'component/pm/task/dialog';
import { getStyles } from './style';
import { roles } from './role';
import { urls, taskUrls, messageUrls } from '../config';

const defaultConversationsItems = [
    {
        key: '0',
        label: 'Default'
    }
];

const useStyle = getStyles(createStyles);

const itemToConversation = (item) => ({
    key: item.id,
    label: item.title,
    description: item.description
});

const conversationToItem = (conversation) => ({
    id: conversation.key,
    title: conversation.label,
    description: conversation.description
});

export default function Chat({ defaultTask, onNav }) {
    const { notification } = App.useApp();
    const { project_id, task_id } = useParams();
    const channel = `${project_id}/${task_id}`;
    const userId = StorageUtil.getUserId();
    const taskId = parseInt(task_id);
    const projectId = parseInt(project_id);
    const navigate = useNavigate();
    const [task, setTask] = useState(defaultTask);
    const [conn, setConn] = useState(null);
    const [isRequesting, setIsRequesting] = useState(false);
    const [messages, setMessages] = useState([]);
    const { styles } = useStyle();

    // ==================== State ====================
    const [headerOpen, setHeaderOpen] = React.useState(false);
    const [content, setContent] = React.useState('');
    const [conversationsItems, setConversationsItems] = React.useState(
        defaultConversationsItems
    );
    const [activeKey, setActiveKey] = React.useState(taskId);
    const [attachedFiles, setAttachedFiles] = React.useState([]);

    const navigateTo = NavUtil.navigateTo(navigate);
    // ==================== Runtime ====================

    useEffect(() => {
        if (activeKey !== undefined) {
            setMessages([]);
        }
        const conversation = conversationsItems.find((item) => item.key === activeKey);
        if (conversation) {
            const item = conversationToItem(conversation);
            handleTaskChange(item);
        }
    }, [activeKey]);

    useEffect(() => {
        getTaskList(taskId);
        getMessage(taskId);
    }, [taskId]);

    const getMessage = (taskId) => {
        return RequestUtil.apiCall(messageUrls.crud, { task_id: taskId })
            .then((resp) => {
                setMessages(resp.data);
            })
            .catch(RequestUtil.displayError(notification));
    };

    const handleTaskChange = (item) => {
        if (!item) {
            return;
        }
        const conversationIndex = conversationsItems.findIndex(
            (conversation) => conversation.key === item.id
        );
        if (conversationIndex !== -1) {
            const conversation = itemToConversation(item);
            conversationsItems[conversationIndex] = conversation;
            setConversationsItems([...conversationsItems]);
        }
        onNav(item.title);
        setTask({
            id: item.id,
            title: item.title,
            description: item.description
        });
    };

    const getTaskList = () => {
        RequestUtil.apiCall(taskUrls.crud, { project_id: projectId })
            .then((resp) => {
                setConversationsItems(resp.data.map(itemToConversation));
            })
            .catch(RequestUtil.displayError(notification));
    };

    useEffect(() => {
        SocketUtil.newConn()
            .then((conn) => {
                setConn(conn);
            })
            .catch(RequestUtil.displayError(notification));
        return () => {
            conn && conn.disconnect();
        };
    }, []);

    useEffect(() => {
        if (!conn) {
            return;
        }
        handleConnect(conn);
        const sub = handleSubscription(conn, channel);

        return () => {
            if (sub && sub.state === 'subscribed' && conn) {
                sub.unsubscribe();
                sub.removeAllListeners();
                conn.removeSubscription(sub);
            }
        };
    }, [conn, channel]);

    const handleConnect = (conn) => {
        // Event Handlers
        conn.on('connecting', (ctx) => {
            console.log(`connecting: ${ctx.code}, ${ctx.reason}`);
        });

        conn.on('connected', (ctx) => {
            console.log('connected', ctx);
        });

        conn.on('disconnected', (ctx) => {
            console.log(`disconnected: ${ctx.code}, ${ctx.reason}`);
        });

        if (conn.state === 'connected') {
            return conn;
        }
        conn.connect();

        return conn;
    };

    const handleSubscription = (conn, channel) => {
        // Subscribe to the channel
        const existSub = conn.getSubscription(channel);
        if (existSub) {
            return existSub;
        }
        const sub = conn.newSubscription(channel);

        sub.on('publication', (ctx) => {
            const { data } = ctx;
            console.log('publication', data);
            data.editable = data.user_id === userId;
            handleAddMessage(data);
        });
        /*
        sub.on('subscribing', (ctx) => {
            console.log(`subscribing: ${ctx.code}, ${ctx.reason}`);
        });
        */
        sub.on('subscribed', (ctx) => {
            console.log('subscribed', ctx);
        });
        sub.on('unsubscribed', (ctx) => {
            console.log(`unsubscribed: ${ctx.code}, ${ctx.reason}`);
        });

        // Subscribe to the channel
        sub.subscribe();
        return sub;
    };

    const handleChange = (data, id) => {
        const item = { id, title: data.title, description: data.description };
        setTask(item);
        handleTaskChange(item);
        /*
        if (!id) {
            setList([{ ...Util.appendKey(data) }, ...list]);
        } else {
            setTask(data);
        }
        */
    };

    const handleDelete = (id) => {
        const r = window.confirm('Do you want to remove this task?');
        if (!r) {
            return;
        }

        Util.toggleGlobalLoading(true);
        RequestUtil.apiCall(`${urls.crud}${id}`, {}, 'delete')
            .then(() => {
                Dialog.toggle(false);
                navigateTo(`/pm/task/${projectId}`);
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                Util.toggleGlobalLoading(false);
            });
    };

    // ==================== Event ====================

    const publishMessage = (content) => {
        setIsRequesting(true);
        const payload = {
            channel,
            project_id: projectId,
            task_id: taskId,
            content
        };
        attachedFiles.forEach((file, index) => {
            payload[`_file_${index}`] = file.originFileObj;
        });
        return RequestUtil.apiCall(urls.publishMessage, payload, 'post')
            .then((resp) => {
                return resp;
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                setIsRequesting(false);
            });
    };

    const handleAddMessage = (message) => {
        setMessages((messages) => [...messages, { ...message }]);
    };

    const handleSending = (nextContent) => {
        if (!nextContent) {
            return;
        }
        setContent('');
        publishMessage(nextContent);
    };

    const onConversationClick = (key) => {
        setActiveKey(key);
        navigateTo(`/pm/task/${projectId}/${key}`);
    };
    const handleFileChange = (info) => {
        console.log('File change:', info);
        console.log(typeof info.file, info.file instanceof Blob);
        console.log(
            typeof info.fileList[0],
            info.fileList[0].originFileObj instanceof Blob
        );
        setAttachedFiles(info.fileList);
    };

    // ==================== Nodes ====================
    const items = messages.map((message) => {
        const editable = message.user.id === userId;
        message.editable = editable;
        return message;
    });
    const attachmentsNode = (
        <Badge dot={attachedFiles.length > 0 && !headerOpen}>
            <Button
                type="text"
                icon={<PaperClipOutlined />}
                onClick={() => {
                    setHeaderOpen(!headerOpen);
                }}
            />
        </Badge>
    );
    const senderHeader = (
        <Sender.Header
            title="Attachments"
            open={headerOpen}
            onOpenChange={setHeaderOpen}
            styles={{
                content: {
                    padding: 0
                }
            }}
        >
            <Attachments
                beforeUpload={() => false}
                items={attachedFiles}
                onChange={handleFileChange}
                placeholder={(type) =>
                    type === 'drop'
                        ? {
                              title: 'Drop file here'
                          }
                        : {
                              icon: <CloudUploadOutlined />,
                              title: 'Upload files',
                              description: 'Click or drag files to this area to upload'
                          }
                }
            />
        </Sender.Header>
    );

    // ==================== Render =================
    return (
        <>
            <div className={styles.layout}>
                <div className={styles.menu}>
                    <Conversations
                        items={conversationsItems}
                        className={styles.conversations}
                        activeKey={activeKey}
                        onActiveChange={onConversationClick}
                    />
                </div>
                <div className={styles.chat}>
                    <div className="flex-container">
                        <div className="flex-item-remaining">
                            <div>
                                <strong># {task.title}</strong>
                            </div>
                            <div>{task.description}</div>
                        </div>
                        <div>
                            <Button
                                onClick={() => TaskDialog.toggle(true, task.id)}
                                icon={<EditOutlined />}
                            />
                        </div>
                    </div>
                    {/*
                    <Bubble.List
                        items={items}
                        roles={roles}
                        className={styles.messages}
                    />
                    */}
                    <Flex gap="middle" vertical>
                        {items.map((item) => {
                            return (
                                <Bubble
                                    key={item.id}
                                    content={item.content}
                                    className={styles.message}
                                    header={item.user.name}
                                    avatar={{
                                        icon: (
                                            <Avatar
                                                size="small"
                                                src={item.user.avatar}
                                            />
                                        )
                                    }}
                                />
                            );
                        })}
                    </Flex>
                    <Sender
                        value={content}
                        header={senderHeader}
                        onSubmit={handleSending}
                        onChange={setContent}
                        prefix={attachmentsNode}
                        loading={isRequesting}
                        className={styles.sender}
                    />
                </div>
            </div>
            <TaskDialog onChange={handleChange} onDelete={handleDelete} />
        </>
    );
}

Chat.displayName = 'Chat';
