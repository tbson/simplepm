import * as React from 'react';
import { useEffect, useState } from 'react';
import { App, Badge, Button, Space } from 'antd';
import {
    Attachments,
    Bubble,
    Conversations,
    Sender,
    Welcome,
    useXAgent,
    useXChat
} from '@ant-design/x';
import { createStyles } from 'antd-style';
import { CloudUploadOutlined, PaperClipOutlined } from '@ant-design/icons';
import { Centrifuge } from 'centrifuge';
import RequestUtil from 'service/helper/request_util';
import {
    CENTRIFUGO_SUBSCRIPTION_TOKEN_ENDPOINT,
    CENTRIFUGO_SOCKET_ENDPOINT
} from 'src/const';
import { getStyles } from './style';
import { roles } from './role';
import { urls } from '../config';

const defaultConversationsItems = [
    {
        key: '0',
        label: 'Default'
    },
    {
        key: '1',
        label: 'Feature 1'
    }
];

const useStyle = getStyles(createStyles);

async function getToken(ctx) {
    const res = await fetch(CENTRIFUGO_SUBSCRIPTION_TOKEN_ENDPOINT, {
        method: 'POST',
        headers: new Headers({ 'Content-Type': 'application/json' }),
        body: JSON.stringify({
            channel: ctx.channel
        })
    });
    if (!res.ok) {
        if (res.status === 403) {
            // Return special error to not proceed with token refreshes,
            // client will be disconnected.
            throw new Centrifuge.UnauthorizedError();
        }
        // Any other error thrown will result into token refresh re-attempts.
        throw new Error(`Unexpected status code ${res.status}`);
    }
    const data = await res.json();
    return data.token;
}

export default function Chat() {
    const { notification } = App.useApp();
    const [token, setToken] = useState('');
    const [count, setCount] = useState('-');
    const [connectionStatus, setConnectionStatus] = useState('Disconnected');

    const { styles } = useStyle();

    // ==================== State ====================
    const [headerOpen, setHeaderOpen] = React.useState(false);
    const [content, setContent] = React.useState('');
    const [conversationsItems, setConversationsItems] = React.useState(
        defaultConversationsItems
    );
    const [activeKey, setActiveKey] = React.useState(defaultConversationsItems[0].key);
    const [attachedFiles, setAttachedFiles] = React.useState([]);

    // ==================== Runtime ====================
    const [agent] = useXAgent({
        request: async ({ message }, { onSuccess }) => {
            onSuccess(`Mock success return. You said: ${message}`);
        }
    });
    const { onRequest, messages, setMessages } = useXChat({
        agent
    });
    useEffect(() => {
        if (activeKey !== undefined) {
            setMessages([]);
        }
    }, [activeKey]);

    useEffect(() => {
        RequestUtil.apiCall(urls.getJwt)
            .then((resp) => {
                setToken(resp.data.token);
            })
            .catch(RequestUtil.displayError(notification));
    }, []);

    useEffect(() => {
        if (!token) return;
        const centrifuge = new Centrifuge(CENTRIFUGO_SOCKET_ENDPOINT, {
            token,
            getToken
        });

        // Event Handlers
        centrifuge.on('connecting', (ctx) => {
            console.log(`connecting: ${ctx.code}, ${ctx.reason}`);
            setConnectionStatus(`Connecting (${ctx.code}): ${ctx.reason}`);
        });

        centrifuge.on('connected', (ctx) => {
            console.log(`connected over ${ctx.transport}`);
            setConnectionStatus(`Connected via ${ctx.transport}`);
        });

        centrifuge.on('disconnected', (ctx) => {
            console.log(`disconnected: ${ctx.code}, ${ctx.reason}`);
            setConnectionStatus(`Disconnected (${ctx.code}): ${ctx.reason}`);
        });

        // Connect to Centrifugo
        centrifuge.connect();

        // Subscribe to the channel
        const sub = centrifuge.newSubscription('channel');

        sub.on('publication', (ctx) => {
            if (ctx.data && typeof ctx.data.value !== 'undefined') {
                setCount(ctx.data.value);
                document.title = ctx.data.value.toString();
            }
        });

        sub.on('subscribing', (ctx) => {
            console.log(`subscribing: ${ctx.code}, ${ctx.reason}`);
        });

        sub.on('subscribed', (ctx) => {
            console.log('subscribed', ctx);
        });

        sub.on('unsubscribed', (ctx) => {
            console.log(`unsubscribed: ${ctx.code}, ${ctx.reason}`);
        });

        // Subscribe to the channel
        sub.subscribe();

        // Cleanup function to disconnect Centrifuge when the component unmounts
        return () => {
            sub.unsubscribe();
            centrifuge.disconnect();
        };
    }, [token]);

    // ==================== Event ====================
    const onSubmit = (nextContent) => {
        if (!nextContent) return;
        onRequest(nextContent);
        setContent('');
    };
    const onPromptsItemClick = (info) => {
        onRequest(info.data.description);
    };
    const onAddConversation = () => {
        setConversationsItems([
            ...conversationsItems,
            {
                key: `${conversationsItems.length}`,
                label: `New Conversation ${conversationsItems.length}`
            }
        ]);
        setActiveKey(`${conversationsItems.length}`);
    };
    const onConversationClick = (key) => {
        setActiveKey(key);
    };
    const handleFileChange = (info) => setAttachedFiles(info.fileList);

    // ==================== Nodes ====================
    const placeholderNode = (
        <Space direction="vertical" size={16} className={styles.placeholder}>
            <Welcome
                variant="borderless"
                title="#Feature 1"
                description="Feature 1 description"
            />
        </Space>
    );
    const items = messages.map(({ id, message, status }) => ({
        key: id,
        loading: status === 'loading',
        role: status === 'you' ? 'you' : 'their',
        content: message
    }));
    const attachmentsNode = (
        <Badge dot={attachedFiles.length > 0 && !headerOpen}>
            <Button
                type="text"
                icon={<PaperClipOutlined />}
                onClick={() => setHeaderOpen(!headerOpen)}
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
                    <Bubble.List
                        items={
                            items.length > 0
                                ? items
                                : [
                                      {
                                          content: placeholderNode,
                                          variant: 'borderless'
                                      }
                                  ]
                        }
                        roles={roles}
                        className={styles.messages}
                    />
                    <Sender
                        value={content}
                        header={senderHeader}
                        onSubmit={onSubmit}
                        onChange={setContent}
                        prefix={attachmentsNode}
                        loading={agent.isRequesting()}
                        className={styles.sender}
                    />
                </div>
            </div>
        </>
    );

    /*
    const messages = getMessages();
    return (
        <>
            <PageHeading>
                <>{messages.heading}</>
            </PageHeading>
            <div style={styles.container}>
                <h1>Centrifugo Counter</h1>
                <div style={styles.counter}>{count}</div>
                <div style={styles.status}>Status: {connectionStatus}</div>
            </div>
        </>
    );
    */
}
/*
const styles = {
    container: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        fontFamily: 'Arial, sans-serif'
    },
    counter: {
        fontSize: '48px',
        margin: '20px 0'
    },
    status: {
        fontSize: '16px',
        color: '#555'
    }
};
*/

Chat.displayName = 'Chat';
