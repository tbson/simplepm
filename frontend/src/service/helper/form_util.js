import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import DateUtil from 'service/helper/date_util';

export default class FormUtil {
    /**
     * removeEmptyKey.
     *
     * @param {Object} form - Antd hook instance
     * @param {Object} errorDict - {str: str[]}
     */

    static setFormErrors(form = null, notification = null) {
        return ({ errors }) => {
            const errorDict = {};
            for (const error of errors) {
                errorDict[error.field] = error.messages;
            }
            if ('detail' in errorDict && notification) {
                notification.error({
                    message: 'Error',
                    description: errorDict.detail,
                    duration: 5
                });
                delete errorDict.detail;
            }
            form &&
                form.setFields(
                    Object.entries(errorDict).map(([name, errors]) => ({
                        name,
                        errors: typeof errors === 'string' ? [errors] : errors
                    }))
                );
        };
    }

    /**
     * handleSubmit.
     *
     * @param {Object} payload
     */
    static submit(url, payload, method = 'post') {
        Util.toggleGlobalLoading();
        return new Promise((resolve, reject) => {
            RequestUtil.apiCall(url, payload, method)
                .then((resp) => {
                    resolve(resp.data);
                })
                .catch((err) => {
                    reject(err.response.data);
                })
                .finally(() => Util.toggleGlobalLoading(false));
        });
    }

    /**
     * getDefaultFieldName.
     *
     * @param {String} fieldName
     * @returns {String}
     */
    static getDefaultFieldName(fieldName) {
        return fieldName ? `"${fieldName}"` : 'này';
    }

    /**
     * ruleRequired.
     *
     * @param {String} fieldName
     * @returns {Object} - Antd Form Rule Object
     */
    static ruleRequired(fieldName = '') {
        fieldName = FormUtil.getDefaultFieldName(fieldName);
        return {
            required: true,
            message: `Trường ${fieldName} là bắt buộc`
        };
    }

    /**
     * ruleMin.
     *
     * @param {Number} min
     * @param {String} fieldName
     * @returns {Object} - Antd Form Rule Object
     */
    static ruleMin(min, fieldName = '') {
        fieldName = FormUtil.getDefaultFieldName(fieldName);
        return {
            type: 'number',
            min,
            message: `Trường "${fieldName}" có giá trị bé nhất là: ${min}`
        };
    }

    /**
     * ruleMax.
     *
     * @param {Number} max
     * @param {String} fieldName
     * @returns {Object} - Antd Form Rule Object
     */
    static ruleMax(max, fieldName = '') {
        fieldName = FormUtil.getDefaultFieldName(fieldName);
        return {
            type: 'number',
            max,
            message: `Trường "${fieldName}" có giá trị lớn nhất là: ${max}`
        };
    }

    static addOptional(options, value = null, label = '--Select--') {
        return [{ value, label }, ...options];
    }

    static parseFieldValue(value, type) {
        if (type === 'DATE') {
            if (!value) {
                return null;
            }
            return DateUtil.strToDate(value);
        }
        if (type === 'MULTIPLE_SELECT') {
            const result = value.split(',').map((v) => parseInt(v));
            if (result.length === 1 && !result[0]) {
                return [];
            }
            return result;
        }
        if (['NUMBER', 'SELECT'].includes(type)) {
            if (!value) {
                return null;
            }
            return parseInt(value);
        }
        return value;
    }

    static getFile() {
        return new Promise((resolve) => {
            const input = document.createElement('input');
            input.type = 'file';
            input.style.display = 'none';

            document.body.appendChild(input);

            input.addEventListener('change', (event) => {
                const file = event.target.files[0];
                if (file) {
                    resolve(file);
                }
                document.body.removeChild(input);
            });

            input.click();
        });
    }
}
