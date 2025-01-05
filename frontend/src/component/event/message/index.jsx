import * as React from 'react';
import { useEffect, useState } from 'react';
import { App } from 'antd';
import { Centrifuge } from 'centrifuge';
import RequestUtil from 'service/helper/request_util';
import PageHeading from 'component/common/page_heading';
import {
    CENTRIFUGO_SUBSCRIPTION_TOKEN_ENDPOINT,
    CENTRIFUGO_SOCKET_ENDPOINT
} from 'src/const';
import { urls, getMessages } from './config';

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
}

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

Message.displayName = 'Message';
