import * as React from 'react';
import { useEffect, useState, useRef, useCallback } from 'react';
import { useNavigate } from 'react-router';
import { useSetAtom } from 'jotai';
import { App, Badge, Button, Flex, Avatar, Dropdown, Space } from 'antd';
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
    ArrowUpOutlined,
    MenuOutlined
} from '@ant-design/icons';
import Util from 'service/helper/util';
import NavUtil from 'service/helper/nav_util';
import RequestUtil from 'service/helper/request_util';
import SocketUtil from 'service/helper/socket_util';
import StorageUtil from 'service/helper/storage_util';
import { taskOptionSt } from 'component/pm/task/state';
import TaskDialog from 'component/pm/task/dialog';
import DocTable from 'component/document/doc/table';
import { getStyles } from './style';
import { roles } from './role';
import { urls, taskUrls } from '../config';

const START_INDEX = 100000;
const MESSAGE_CREATED = 'MESSAGE_CREATED';
const MESSAGE_UPDATED = 'MESSAGE_UPDATED';
const MESSAGE_DELETED = 'MESSAGE_DELETED';
const GIT_PUSHED = 'GIT_PUSHED';

const useStyle = getStyles(createStyles);

const itemToConversation = (item) => ({
    key: item.id,
    label: item.title,
    description: item.description
});

export default function Chat({ project, defaultTask, onNav }) {
    const { notification } = App.useApp();
    const userId = StorageUtil.getUserId();
    const setTaskOption = useSetAtom(taskOptionSt);
    const { id: projectId, title: projectTitle } = project;
    const [taskId, setTaskId] = useState(defaultTask.id);
    const channel = `${projectId}/${taskId}`;
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
    const [showDocList, setShowDocList] = useState(true);
    const { styles } = useStyle();

    // ==================== State ====================
    const [headerOpen, setHeaderOpen] = React.useState(false);
    const [content, setContent] = React.useState('');
    const [taskList, setTaskList] = React.useState([]);
    const [attachedFiles, setAttachedFiles] = React.useState([]);

    const navigateTo = NavUtil.navigateTo(navigate);
    // ==================== Runtime ====================

    useEffect(() => {
        if (!taskId) return;
        getOption(projectId);
        getTaskList()
            .then(handleTaskChange)
            .then(() => getMessage(true));
    }, [taskId]);

    const getOption = (projectId) => {
        RequestUtil.apiCall(taskUrls.option, { project_id: projectId })
            .then((resp) => {
                setTaskOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setTaskOption((prev) => ({ ...prev, loaded: true }));
            });
    };

    const getTaskList = () => {
        return RequestUtil.apiCall(taskUrls.crud, { project_id: projectId })
            .then((resp) => {
                const taskList = resp.data;
                setTaskList(taskList);
                return taskList;
            })
            .catch(RequestUtil.displayError(notification));
    };

    const handleTaskChange = (taskList) => {
        let item = {};
        const index = taskList.findIndex((item) => item.id === taskId);
        if (index !== -1) {
            item = taskList[index];
            taskList[index] = item;
            setTaskList([...taskList]);
        }
        onNav(item.title);
        setTask({
            id: item.id,
            title: item.title,
            description: item.description
        });
    };

    const handleSingleTaskChange = () => {
        getTaskList()
            .then(handleTaskChange)
            .then(() => getMessage(true));
    };

    const getMessage = (isInit = false) => {
        const params = {
            task_id: taskId
        };
        if (pageState && !isInit) {
            params.page_state = pageState;
        }
        return RequestUtil.apiCall(urls.crud, params)
            .then((resp) => {
                const newMsgs = resp.data.items;
                setMessages((messages) => {
                    const finalMessages = isInit ? newMsgs : [...newMsgs, ...messages];
                    return formatMessages(finalMessages);
                });
                setPageState(resp.data.page_state);
                setFirstItemIndex((index) => {
                    return (isInit ? START_INDEX : index) - newMsgs.length;
                });
            })
            .catch(RequestUtil.displayError(notification));
    };

    useEffect(() => {
        let connection; // local variable to keep track of the connection

        SocketUtil.newConn()
            .then((conn) => {
                connection = conn;
                // Register event listeners once
                conn.on('connecting', (ctx) => {
                    // console.log(`Connecting: ${ctx.code}, ${ctx.reason}`);
                });
                conn.on('connected', (ctx) => {
                    // console.log('Connected', ctx);
                });
                conn.on('disconnected', (ctx) => {
                    // console.log(`Disconnected: ${ctx.code}, ${ctx.reason}`);
                });
                // Ensure the connection is active
                if (conn.state !== 'connected') {
                    conn.connect();
                }
                setConn(conn);
            })
            .catch(RequestUtil.displayError(notification));

        // Cleanup on unmount: disconnect the connection
        return () => {
            if (connection) {
                connection.disconnect();
            }
        };
    }, []);

    useEffect(() => {
        if (!conn) {
            return;
        }
        const sub = handleSubscription(conn, channel);

        return () => {
            if (sub && sub.state === 'subscribed' && conn) {
                sub.unsubscribe();
                sub.removeAllListeners();
                conn.removeSubscription(sub);
            }
        };
    }, [conn, channel]);

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
            if ([MESSAGE_CREATED, GIT_PUSHED].includes(data.type)) {
                handleAddMessage(data);
            }
            if (data.type === MESSAGE_UPDATED) {
                handleUpdateMessage(data);
            }
            if (data.type === MESSAGE_DELETED) {
                handleDeleteMessage(data);
            }
        });
        /*
        sub.on('subscribing', (ctx) => {
            console.log(`subscribing: ${ctx.code}, ${ctx.reason}`);
        });
        */
        sub.on('subscribed', (ctx) => {
            // console.log('subscribed', ctx);
        });
        sub.on('unsubscribed', (ctx) => {
            // console.log(`unsubscribed: ${ctx.code}, ${ctx.reason}`);
        });

        // Subscribe to the channel
        sub.subscribe();
        return sub;
    };

    const handleChange = (data, id) => {
        const item = { id, title: data.title, description: data.description };
        setTask(item);
        handleSingleTaskChange();
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
        RequestUtil.apiCall(`${taskUrls.crud}${id}`, {}, 'delete')
            .then(() => {
                TaskDialog.toggle(false);
                navigateTo(`/pm/project/${projectId}`);
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
        setTaskId(key);
        navigateTo(`/pm/task/${key}`);
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
        return renderAttachments(item.attachments || []);
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

    const handleStartReached = () => {
        if (pageState === null) {
            return;
        }
        getMessage();
    };

    const renderMessage = (item) => {
        return <Markdown>{item.content}</Markdown>;
    };

    const renderGitPushed = (item) => {
        return (
            <div>
                <em>
                    Git pushed: to branch: <strong>{item.git_data.git_branch}</strong>
                </em>
                <ul>
                    {item.git_data.git_commits.map((commit, index) => (
                        <li key={index}>
                            <a href={commit.commit_url} target="_blank">
                                {commit.commit_message}
                            </a>
                        </li>
                    ))}
                </ul>
            </div>
        );
    };

    const renderGitPr = (item) => {
        return "PR...";
        return (
            <div>
                <em>
                    Git pushed: to branch: <strong>{item.git_data.git_branch}</strong>
                </em>
                <ul>
                    {item.git_data.git_commits.map((commit, index) => (
                        <li key={index}>
                            <a href={commit.commit_url} target="_blank">
                                {commit.commit_message}
                            </a>
                        </li>
                    ))}
                </ul>
            </div>
        );
    };

    // ==================== Render =================
    return (
        <>
            <div className={styles.layout}>
                <div className={styles.menu}>
                    <div className={styles.chatHeading}>{projectTitle}</div>
                    <Conversations
                        items={taskList.map(itemToConversation)}
                        className={styles.conversations}
                        activeKey={taskId}
                        onActiveChange={onConversationClick}
                    />
                </div>
                <div className={styles.chat}>
                    <div className={styles.chatHeading} style={{ paddingRight: 0 }}>
                        <div className="flex-item-remaining">
                            <div>
                                <strong># {task.title}</strong>
                            </div>
                            <div>{task.description}</div>
                        </div>
                        <div>
                            <Button
                                danger
                                onClick={() => handleDelete(task.id)}
                                icon={<DeleteOutlined />}
                            />
                            &nbsp;
                            <Button
                                onClick={() => TaskDialog.toggle(true, task.id)}
                                icon={<EditOutlined />}
                            />
                            &nbsp;
                            <Button
                                color="default"
                                variant="link"
                                onClick={() => setShowDocList(!showDocList)}
                                icon={<MenuOutlined />}
                            />
                        </div>
                    </div>
                    <Virtuoso
                        key={taskId}
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
                                    content={item}
                                    messageRender={(item) => {
                                        if (item.type === 'MESSAGE_CREATED') {
                                            return renderMessage(item);
                                        }
                                        if (item.type === 'GIT_PUSHED') {
                                            return renderGitPushed(item);
                                        }
                                        if (item.type === 'GIT_PR_CREATED') {
                                            return renderGitPr(item);
                                        }
                                        return <div>---INVALID TYPE---</div>;
                                    }}
                                    className={styles.messages}
                                    header={item.user.name}
                                    avatar={
                                        <Flex gap="middle" vertical>
                                            <Avatar
                                                size="small"
                                                shape="square"
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
                {showDocList ? <DocTable taskId={taskId} showControl /> : null}
            </div>
            <TaskDialog projectId={projectId} onChange={handleChange} />
        </>
    );
}

Chat.displayName = 'Chat';
