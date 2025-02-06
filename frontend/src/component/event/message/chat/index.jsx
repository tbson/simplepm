import * as React from 'react';
import { useEffect, useState, useRef, useCallback } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { App, Badge, Button, Flex, Avatar, Dropdown, Space, List } from 'antd';
import { Attachments, Bubble, Conversations, Sender } from '@ant-design/x';
import { Virtuoso } from 'react-virtuoso';
import Markdown from 'react-markdown';
import { createStyles } from 'antd-style';
import {
    CloudUploadOutlined,
    PaperClipOutlined,
    EditOutlined,
    DeleteOutlined,
    MoreOutlined,
    ArrowUpOutlined
} from '@ant-design/icons';
import Util from 'service/helper/util';
import NavUtil from 'service/helper/nav_util';
import RequestUtil from 'service/helper/request_util';
import SocketUtil from 'service/helper/socket_util';
import StorageUtil from 'service/helper/storage_util';
import TaskDialog from 'component/pm/task/dialog';
import Doc from 'component/document/doc';
import { getStyles } from './style';
import { roles } from './role';
import { urls, taskUrls } from '../config';

const START_INDEX = 999999;
const CREATE_MESSAGE = 'CREATE_MESSAGE';
const UPDATE_MESSAGE = 'UPDATE_MESSAGE';
const DELETE_MESSAGE = 'DELETE_MESSAGE';

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
    const senderRef = useRef(null);
    const virtuoso = useRef(null);
    const navigate = useNavigate();
    const [firstItemIndex, setFirstItemIndex] = useState(START_INDEX);
    const [task, setTask] = useState(defaultTask);
    const [conn, setConn] = useState(null);
    const [isRequesting, setIsRequesting] = useState(false);
    const [editId, setEditId] = useState(null);
    const [messages, setMessages] = useState([]);
    const [pageState, setPageState] = useState('');
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
        getMessage();
    }, [taskId]);

    const getMessage = () => {
        const params = {
            task_id: taskId
        };
        if (pageState) {
            params.page_state = pageState;
        }
        return RequestUtil.apiCall(urls.crud, params)
            .then((resp) => {
                const newMessages = resp.data.items;
                setMessages((messages) =>
                    formatMessages([...newMessages, ...messages])
                );
                setPageState(resp.data.page_state);
                setFirstItemIndex((index) => index - newMessages.length);
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
            if (data.type === CREATE_MESSAGE) {
                handleAddMessage(data);
            }
            if (data.type === UPDATE_MESSAGE) {
                handleUpdateMessage(data);
            }
            if (data.type === DELETE_MESSAGE) {
                handleDeleteMessage(data);
            }
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

    const createMessage = (content) => {
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
        return RequestUtil.apiCall(urls.crud, payload, 'post')
            .then((resp) => {
                return resp;
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                setIsRequesting(false);
            });
    };

    const updateMessage = (id, content) => {
        setIsRequesting(true);
        const payload = {
            channel,
            content
        };
        return RequestUtil.apiCall(`${urls.crud}${id}/${taskId}`, payload, 'put')
            .then((resp) => {
                return resp;
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                setIsRequesting(false);
                setEditId(null);
            });
    };

    const deleteMessage = (id) => {
        setIsRequesting(true);
        const payload = {
            channel
        };
        return RequestUtil.apiCall(`${urls.delete}${id}/${taskId}`, payload, 'put')
            .then((resp) => {
                return resp;
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                setIsRequesting(false);
            });
    };

    const handleAddMessage = (message) => {
        setMessages((messages) => {
            setTimeout(() => {
                virtuoso.current.scrollToIndex({
                    index: messages.length,
                    align: 'end',
                    behavior: 'smooth'
                });
            }, 250);

            return formatMessages([...messages, { ...message }]);
        });
        setFirstItemIndex((index) => index - 1);
    };

    const handleUpdateMessage = (data) => {
        setMessages((messages) => {
            const index = messages.findIndex((item) => item.id === data.id);
            if (index === -1) {
                return messages;
            }
            const newMessages = [...messages];
            newMessages[index].content = data.content;
            return newMessages;
        });
    };

    const handleDeleteMessage = (data) => {
        setMessages((messages) => messages.filter((item) => item.id !== data.id));
    };

    const handleSending = (nextContent) => {
        if (!nextContent) {
            return;
        }
        setContent('');
        if (editId) {
            updateMessage(editId, nextContent);
        } else {
            createMessage(nextContent);
        }
        setAttachedFiles([]);
        setHeaderOpen(false);
    };

    const onConversationClick = (key) => {
        setActiveKey(key);
        navigateTo(`/pm/task/${projectId}/${key}`);
    };
    const handleFileChange = (info) => {
        setAttachedFiles(info.fileList);
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
                        const r = window.confirm('Do you want to remove this message?');
                        if (!r) {
                            return;
                        }
                        deleteMessage(item.id);
                    }
                }
            ]
        };
    };

    // ==================== Nodes ====================
    const formatMessages = (messages) => {
        return messages.map((message) => {
            const editable = message.user.id === userId;
            message.editable = editable;
            return message;
        });
    };
    const attachmentsNode = (
        <Badge dot={attachedFiles.length > 0 && !headerOpen}>
            <Button
                icon={<PaperClipOutlined />}
                onClick={() => {
                    setHeaderOpen(!headerOpen);
                    // scroll main content to bottom
                    setTimeout(() => {
                        window.scrollTo({
                            left: 0,
                            top: document.body.scrollHeight,
                            behavior: 'smooth'
                        });
                    }, 250);
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
                getDropContainer={() => senderRef.current?.nativeElement}
            />
        </Sender.Header>
    );

    const renderFooter = (item) => {
        return renderAttachments(item.attachments);
    };

    const renderAttachments = (files) => {
        const fileBlock = files.map((attachment, index) => {
            const item = {
                uid: index,
                name: attachment.file_name,
                file_url: attachment.file_url,
                size: attachment.file_size
            };
            if (
                attachment.file_type.startsWith('image') &&
                !attachment.file_type.includes('svg')
            ) {
                item.url = attachment.file_url;
            }
            return (
                <a key={index} href={item.file_url} target="_blank">
                    <Attachments.FileCard key={index} item={item} className="pointer" />
                </a>
            );
        });
        return (
            <div>
                {files.length > 0 ? (
                    <Flex gap="middle" align="start">
                        {fileBlock}
                    </Flex>
                ) : null}
            </div>
        );
    };

    const ScrollHeader = () => {
        return (
            <div
                style={{
                    padding: '2rem',
                    display: 'flex',
                    justifyContent: 'center'
                }}
            >
                Loading...
            </div>
        );
    };

    const handleStartReached = useCallback(() => {
        if (pageState === null) {
            return;
        }
        getMessage();
    }, [pageState]);

    // ==================== Render =================
    return (
        <>
            <div className={styles.layout}>
                <div className={styles.menu}>
                    <div className={styles.chatHeading}>Project name...</div>
                    <Conversations
                        items={conversationsItems}
                        className={styles.conversations}
                        activeKey={activeKey}
                        onActiveChange={onConversationClick}
                    />
                </div>
                <div className={styles.chat}>
                    <div className={styles.chatHeading}>
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
                    <Virtuoso
                        ref={virtuoso}
                        initialTopMostItemIndex={firstItemIndex}
                        firstItemIndex={firstItemIndex}
                        style={{
                            height: 'calc(100vh - 250px)',
                            paddingLeft: 5,
                            paddingRight: 5,
                            overflowX: 'hidden'
                        }}
                        data={messages}
                        startReached={handleStartReached}
                        itemContent={(_index, item) => {
                            return (
                                <Bubble
                                    key={item.id}
                                    content={item.content}
                                    messageRender={(content) => (
                                        <Markdown>{content}</Markdown>
                                    )}
                                    className={styles.messages}
                                    header={item.user.name}
                                    avatar={
                                        <Flex gap="middle" vertical>
                                            <Avatar
                                                size="small"
                                                src={item.user.avatar}
                                            />
                                            {item.editable ? (
                                                <Dropdown
                                                    menu={getMessageMenuItems(item)}
                                                    trigger={['click']}
                                                >
                                                    <Button
                                                        style={{ top: -4 }}
                                                        color="default"
                                                        variant="link"
                                                        icon={
                                                            <MoreOutlined
                                                                style={{
                                                                    fontSize: '20px'
                                                                }}
                                                            />
                                                        }
                                                    />
                                                </Dropdown>
                                            ) : null}
                                        </Flex>
                                    }
                                    footer={renderFooter(item)}
                                />
                            );
                        }}
                        components={{
                            Header: pageState !== null ? ScrollHeader : null
                        }}
                    />

                    <Sender
                        alignToBottom
                        ref={senderRef}
                        value={content}
                        header={senderHeader}
                        onSubmit={handleSending}
                        onChange={setContent}
                        prefix={editId ? null : attachmentsNode}
                        loading={isRequesting}
                        className={styles.sender}
                        actions={(_, info) => {
                            const { SendButton, ClearButton } = info.components;
                            return (
                                <Space size="small">
                                    {editId ? (
                                        <ClearButton onClick={() => setEditId(null)} />
                                    ) : null}
                                    <SendButton
                                        type="primary"
                                        icon={<ArrowUpOutlined />}
                                        disabled={isRequesting}
                                    />
                                </Space>
                            );
                        }}
                    />
                </div>
                <Doc taskId={taskId} />
            </div>
            <TaskDialog onChange={handleChange} onDelete={handleDelete} />
        </>
    );
}

Chat.displayName = 'Chat';
