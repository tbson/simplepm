import { Centrifuge } from 'centrifuge';
import RequestUtil from 'service/helper/request_util';
import {
    CENTRIFUGO_SOCKET_ENDPOINT,
    CENTRIFUGO_SUBSCRIPTION_TOKEN_ENDPOINT
} from 'src/const';

export default class SocketUtil {
    static async getToken(ctx) {
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

    static newConn() {
        if (window.socConn) {
            return Promise.resolve(window.socConn);
        }
        return SocketUtil.getConnToken().then((token) => {
            window.socConn = new Centrifuge(CENTRIFUGO_SOCKET_ENDPOINT, {
                token,
                getToken: SocketUtil.getToken
            });
            return Promise.resolve(window.socConn);
        });
    }

    static getConnToken() {
        const url = '/socket/auth-jwt/';
        return RequestUtil.apiCall(url)
            .then((resp) => {
                return resp.data.token;
            })
            .catch((err) => {
                return Promise.reject(err);
            });
    }
}
