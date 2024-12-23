import * as React from 'react';
import { useEffect, useState } from 'react';
import { Centrifuge } from 'centrifuge';
import PageHeading from 'component/common/page_heading';
import { getMessages } from './config';

export default function Message() {
    const [count, setCount] = useState('-');
    const [connectionStatus, setConnectionStatus] = useState('Disconnected');

    useEffect(() => {
        const centrifuge = new Centrifuge('wss://socketstag.simplepm.io/connection/websocket', {
            token: './exec centrifugo gentoken -u 123722'
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
    }, []);

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
