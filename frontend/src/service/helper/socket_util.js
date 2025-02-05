import { Centrifuge } from 'centrifuge';
import RequestUtil from 'service/helper/request_util';
import { CENTRIFUGO_SOCKET_ENDPOINT } from 'src/const';

export default class SocketUtil {
    static newConn() {
        if (window.socConn) {
            return Promise.resolve(window.socConn);
        }
        return SocketUtil.getConnToken().then((token) => {
            window.socConn = new Centrifuge(CENTRIFUGO_SOCKET_ENDPOINT, {
                token,
                getToken: SocketUtil.getConnToken
            });
            return Promise.resolve(window.socConn);
        });
    }

    static getConnToken() {
        const url = '/socket/jwt/auth/';
        return RequestUtil.apiCall(url)
            .then((resp) => {
                return resp.data.token;
            })
            .catch((err) => {
                // check if the error is 403
                if (err.response.status === 403) {
                    // return special error to not proceed with token refreshes,
                    // client will be disconnected.
                    return Promise.reject(new Centrifuge.UnauthorizedError());
                }
                // Any other error thrown will result into token refresh re-attempts.
                return Promise.reject(
                    new Error(`Unexpected status code ${err.response.status}`)
                );
            });
    }
}
