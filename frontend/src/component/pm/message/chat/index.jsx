import * as React from 'react';
import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
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
import {
    CloudUploadOutlined,
    PaperClipOutlined,
    EditOutlined
} from '@ant-design/icons';
import { Centrifuge } from 'centrifuge';
import Util from 'service/helper/util';
import NavUtil from 'service/helper/nav_util';
import RequestUtil from 'service/helper/request_util';
import {
    CENTRIFUGO_SUBSCRIPTION_TOKEN_ENDPOINT,
    CENTRIFUGO_SOCKET_ENDPOINT
} from 'src/const';
import FeatureDialog from 'component/pm/feature/dialog';
import { getStyles } from './style';
import { roles } from './role';
import { urls, featureUrls } from '../config';

const defaultConversationsItems = [
    {
        key: '0',
        label: 'Default'
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

export default function Chat({ defaultFeature, onNav }) {
    const { notification } = App.useApp();
    const { project_id, feature_id } = useParams();
    const featureId = parseInt(feature_id);
    const projectId = parseInt(project_id);
    const navigate = useNavigate();
    const [feature, setFeature] = useState(defaultFeature);
    const [featureList, setFeatureList] = useState([]);
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
    const [activeKey, setActiveKey] = React.useState(featureId);
    const [attachedFiles, setAttachedFiles] = React.useState([]);

    const navigateTo = NavUtil.navigateTo(navigate);
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
        getFeatureList(featureId);
    }, [featureId]);

    const getFeatureList = () => {
        RequestUtil.apiCall(featureUrls.crud, { project_id: projectId })
            .then((resp) => {
                setConversationsItems(
                    resp.data.map((item) => ({
                        key: item.id,
                        label: item.title,
                        description: item.description
                    }))
                );
            })
            .catch(RequestUtil.displayError(notification));
    };

    /*
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
    */

    const handleChange = (data, id) => {
        if (!id) {
            setList([{ ...Util.appendKey(data) }, ...list]);
        } else {
            setFeature(data);
        }
    };

    const handleDelete = (id) => {
        const r = window.confirm('Do you want to remove this feature?');
        if (!r) return;

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
    const onSubmit = (nextContent) => {
        if (!nextContent) return;
        onRequest(nextContent);
        setContent('');
    };
    const onConversationClick = (key) => {
        const item = conversationsItems.find((item) => item.key === key);
        onNav(item.label);
        setFeature({id: key, title: item.label, description: item.description});
        navigateTo(`/pm/task/message/${projectId}/${key}`);
        setActiveKey(key);
    };
    const handleFileChange = (info) => setAttachedFiles(info.fileList);

    // ==================== Nodes ====================
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
                    <div className="flex-container">
                        <div className="flex-item-remaining">
                            <div>
                                <strong># {feature.title}</strong>
                            </div>
                            <div>{feature.description}</div>
                        </div>
                        <div>
                            <Button
                                onClick={() => FeatureDialog.toggle(true, feature.id)}
                                icon={<EditOutlined />}
                            />
                        </div>
                    </div>
                    <Bubble.List
                        items={items}
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
            <FeatureDialog onChange={handleChange} onDelete={handleDelete} />
        </>
    );
}

Chat.displayName = 'Chat';
