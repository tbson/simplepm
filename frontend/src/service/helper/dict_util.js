class DictUtil {
    /**
     * isEmpty.
     *
     * @param {Object} obj
     * @returns {boolean}
     */
    static isEmpty(obj) {
        if (!obj) {
            return true;
        }
        return Object.keys(obj).length === 0 && obj.constructor === Object;
    }
}

export default DictUtil;
