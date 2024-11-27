export default class StorageUtil {
    /**
     * setStorage.
     *
     * @param {string} key
     * @param {string | Dict} value
     * @returns {void}
     */
    static setStorage(key, value) {
        try {
            localStorage.setItem(key, JSON.stringify(value));
        } catch (error) {
            console.log(error);
        }
    }

    /**
     * setStorageObj.
     *
     * @param {Object} input
     * @returns {void}
     */
    static setStorageObj(input) {
        for (const key in input) {
            const value = input[key];
            StorageUtil.setStorage(key, value);
        }
    }

    /**
     * getStorageObj.
     *
     * @param {string} key
     * @returns {Object}
     */
    static getStorageObj(key) {
        try {
            const value = StorageUtil.parseJson(localStorage.getItem(key));
            if (value && typeof value === 'object') {
                return value;
            }
            return {};
        } catch (error) {
            console.log(error);
            return {};
        }
    }

    /**
     * getStorageStr.
     *
     * @param {string} key
     * @returns {string}
     */
    static getStorageStr(key) {
        try {
            const value = StorageUtil.parseJson(localStorage.getItem(key));
            if (!value || typeof value === 'object') {
                return '';
            }
            return String(value);
        } catch (error) {
            return '';
        }
    }

    /**
     * setToken.
     *
     * @param {string} token
     * @returns {void}
     */
    static setToken(token) {
        const authData = StorageUtil.getStorageObj('auth');
        authData['token'] = token;
        StorageUtil.setStorage('auth', authData);
    }

    /**
     * removeStorage.
     *
     * @param {string} key
     * @returns {void}
     */
    static removeStorage(key) {
        localStorage.removeItem(key);
    }

    /**
     * parseJson.
     *
     * @param {string} input
     * @returns {string}
     */
    static parseJson(input) {
        try {
            return JSON.parse(input);
        } catch (error) {
            return String(input);
        }
    }

    /**
     * getUserInfo.
     *
     * @returns {string}
     */
    static getUserInfo() {
        const { userInfo } = StorageUtil.getStorageObj('auth');
        if (!userInfo?.email) {
            return null;
        }
        return userInfo;
    }

    /**
     * getProfileType.
     *
     * @returns {string}
     */
    static getProfileType() {
        const { userInfo } = StorageUtil.getStorageObj('auth');
        return userInfo.profile_type || 'user';
    }

    /**
     * getPermissions.
     *
     * @returns {string[]}
     */
    static getPermissions() {
        const { pemModulesActionsMap } = StorageUtil.getStorageObj('auth');
        return pemModulesActionsMap || {};
    }

    /**
     * getTenantUid.
     *
     * @returns {string}
     */
    static getTenantUid() {
        const { userInfo } = StorageUtil.getStorageObj('auth');
        let result = userInfo?.tenant_uid || '';
        if (!result) {
            result = StorageUtil.getStorageStr('tenantUid');
        }
        return result;
    }

    static setTenantUid(tenantUid) {
        StorageUtil.setStorage('tenantUid', tenantUid);
    }
}
