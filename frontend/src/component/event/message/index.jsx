import * as React from 'react';
import { useEffect, useState } from 'react';
import { App, Badge, Button, Space } from 'antd';
import {
    Attachments,
    Bubble,
    Conversations,
    Prompts,
    Sender,
    Welcome,
    useXAgent,
    useXChat
} from '@ant-design/x';
import { createStyles } from 'antd-style';
import {
    CloudUploadOutlined,
    CommentOutlined,
    EllipsisOutlined,
    FireOutlined,
    HeartOutlined,
    PaperClipOutlined,
    PlusOutlined,
    ReadOutlined,
    ShareAltOutlined,
    SmileOutlined
} from '@ant-design/icons';
import { Centrifuge } from 'centrifuge';
import RequestUtil from 'service/helper/request_util';
import PageHeading from 'component/common/page_heading';
import {
    CENTRIFUGO_SUBSCRIPTION_TOKEN_ENDPOINT,
    CENTRIFUGO_SOCKET_ENDPOINT
} from 'src/const';
import { urls, getMessages } from './config';

const renderTitle = (icon, title) => (
    <Space align="start">
        {icon}
        <span>{title}</span>
    </Space>
);
const defaultConversationsItems = [
    {
        key: '0',
        label: 'What is Ant Design X?'
    }
];
const useStyle = createStyles(({ token, css }) => {
    return {
        layout: css`
            width: 100%;
            min-width: 1000px;
            height: 722px;
            border-radius: ${token.borderRadius}px;
            display: flex;
            background: ${token.colorBgContainer};
            font-family: AlibabaPuHuiTi, ${token.fontFamily}, sans-serif;

            .ant-prompts {
                color: ${token.colorText};
            }
        `,
        menu: css`
            background: ${token.colorBgLayout}80;
            width: 280px;
            height: 100%;
            display: flex;
            flex-direction: column;
        `,
        conversations: css`
            padding: 0 12px;
            flex: 1;
            overflow-y: auto;
        `,
        chat: css`
            height: 100%;
            width: 100%;
            max-width: 700px;
            margin: 0 auto;
            box-sizing: border-box;
            display: flex;
            flex-direction: column;
            padding: ${token.paddingLG}px;
            gap: 16px;
        `,
        messages: css`
            flex: 1;
        `,
        placeholder: css`
            padding-top: 32px;
        `,
        sender: css`
            box-shadow: ${token.boxShadow};
        `,
        logo: css`
            display: flex;
            height: 72px;
            align-items: center;
            justify-content: start;
            padding: 0 24px;
            box-sizing: border-box;

            img {
                width: 24px;
                height: 24px;
                display: inline-block;
            }

            span {
                display: inline-block;
                margin: 0 8px;
                font-weight: bold;
                color: ${token.colorText};
                font-size: 16px;
            }
        `,
        addBtn: css`
            background: #1677ff0f;
            border: 1px solid #1677ff34;
            width: calc(100% - 24px);
            margin: 0 12px 24px 12px;
        `
    };
});
const placeholderPromptsItems = [
    {
        key: '1',
        label: renderTitle(
            <FireOutlined
                style={{
                    color: '#FF4D4F'
                }}
            />,
            'Hot Topics'
        ),
        description: 'What are you interested in?',
        children: [
            {
                key: '1-1',
                description: `What's new in X?`
            },
            {
                key: '1-2',
                description: `What's AGI?`
            },
            {
                key: '1-3',
                description: `Where is the doc?`
            }
        ]
    },
    {
        key: '2',
        label: renderTitle(
            <ReadOutlined
                style={{
                    color: '#1890FF'
                }}
            />,
            'Design Guide'
        ),
        description: 'How to design a good product?',
        children: [
            {
                key: '2-1',
                icon: <HeartOutlined />,
                description: `Know the well`
            },
            {
                key: '2-2',
                icon: <SmileOutlined />,
                description: `Set the AI role`
            },
            {
                key: '2-3',
                icon: <CommentOutlined />,
                description: `Express the feeling`
            }
        ]
    }
];
const senderPromptsItems = [
    {
        key: '1',
        description: 'Hot Topics',
        icon: (
            <FireOutlined
                style={{
                    color: '#FF4D4F'
                }}
            />
        )
    },
    {
        key: '2',
        description: 'Design Guide',
        icon: (
            <ReadOutlined
                style={{
                    color: '#1890FF'
                }}
            />
        )
    }
];
const roles = {
    ai: {
        placement: 'start',
        typing: {
            step: 5,
            interval: 20
        },
        styles: {
            content: {
                borderRadius: 16
            }
        }
    },
    local: {
        placement: 'end',
        variant: 'shadow'
    }
};

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

export default function Message() {
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
                icon="https://mdn.alipayobjects.com/huamei_iwk9zp/afts/img/A*s5sNRo5LjfQAAAAAAAAAAAAADgCCAQ/fmt.webp"
                title="Hello, I'm Ant Design X"
                description="Base on Ant Design, AGI product interface solution, create a better intelligent vision~"
                extra={
                    <Space>
                        <Button icon={<ShareAltOutlined />} />
                        <Button icon={<EllipsisOutlined />} />
                    </Space>
                }
            />
            <Prompts
                title="Do you want?"
                items={placeholderPromptsItems}
                styles={{
                    list: {
                        width: '100%'
                    },
                    item: {
                        flex: 1
                    }
                }}
                onItemClick={onPromptsItemClick}
            />
        </Space>
    );
    const items = messages.map(({ id, message, status }) => ({
        key: id,
        loading: status === 'loading',
        role: status === 'local' ? 'local' : 'ai',
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
    const logoNode = (
        <div className={styles.logo}>
            <img
                src="https://mdn.alipayobjects.com/huamei_iwk9zp/afts/img/A*eco6RrQhxbMAAAAAAAAAAAAADgCCAQ/original"
                draggable={false}
                alt="logo"
            />
            <span>Ant Design X</span>
        </div>
    );

    // ==================== Render =================
    return (
        <div className={styles.layout}>
            <div className={styles.menu}>
                {/* 🌟 Logo */}
                {logoNode}
                {/* 🌟 添加会话 */}
                <Button
                    onClick={onAddConversation}
                    type="link"
                    className={styles.addBtn}
                    icon={<PlusOutlined />}
                >
                    New Conversation
                </Button>
                {/* 🌟 会话管理 */}
                <Conversations
                    items={conversationsItems}
                    className={styles.conversations}
                    activeKey={activeKey}
                    onActiveChange={onConversationClick}
                />
            </div>
            <div className={styles.chat}>
                {/* 🌟 消息列表 */}
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
                {/* 🌟 提示词 */}
                <Prompts items={senderPromptsItems} onItemClick={onPromptsItemClick} />
                {/* 🌟 输入框 */}
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

Message.displayName = 'Message';
