import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'account/auth',
        endpoints: {
            login: 'login'
        }
    },
    signup: {
        prefix: 'account/signup-tenant',
        endpoints: {
            signup: ''
        }
    }
};

export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const signupUrls = RequestUtil.prefixMapValues(urlMap.signup);

export const TOGGLE_DIALOG_EVENT = 'TOGGLE_RESET_PWD_DIALOG';
const headingTxt = t`Profile`;
export const messages = {
    heading: headingTxt
};
