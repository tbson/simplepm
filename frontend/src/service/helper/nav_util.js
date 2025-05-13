import Util from 'service/helper/util';
import StorageUtil from 'service/helper/storage_util';
import RequestUtil from 'service/helper/request_util';

export default class NavUtil {
    /**
     * navigateTo.
     *
     * @param {Navigate} navigate
     */
    static navigateTo(navigate) {
        return (url = '/') => {
            navigate(url);
        };
    }

    static clearAuthData() {
        StorageUtil.removeStorage('auth');
    }

    /**
     * logout.
     *
     * @param {Navigate} navigate
     */
    static logout(navigate) {
        return () => {
            const baseUrl = RequestUtil.getApiBaseUrl();
            const logoutUrl = `${baseUrl}account/auth/logout/`;
            Util.toggleGlobalLoading();
            RequestUtil.apiCall(logoutUrl, {}, 'post').then(() => {
                NavUtil.cleanAndMoveToLoginPage(navigate);
            }).catch((err) => {
                console.error('Logout failed:', err);
            }).finally(() => {
                NavUtil.cleanAndMoveToLoginPage(navigate);
                Util.toggleGlobalLoading(false);
            });
        };
    }

    /**
     * cleanAndMoveToLoginPage.
     *
     * @param {Navigate} navigate
     * @returns {void}
     */
    static cleanAndMoveToLoginPage(navigate) {
        const currentUrl = window.location.pathname;
        NavUtil.clearAuthData();
        let loginUrl = '/login';
        if (currentUrl) {
            loginUrl = `${loginUrl}?next=${currentUrl}`;
        }
        if (navigate) {
            NavUtil.navigateTo(navigate)(loginUrl);
        } else {
            window.location.href = loginUrl;
        }
    }
}
