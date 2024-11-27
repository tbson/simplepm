import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'account/user',
        endpoints: {
            crud: '',
            option: 'option',
        }
    },
    lock: {
        prefix: 'account/lock-user',
        endpoints: {
            lock: '',
        }
    }
};

export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const lockUrls = RequestUtil.prefixMapValues(urlMap.lock);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_USER_DIALOG';
export const TOGGLE_LOCK_DIALOG_EVENT = 'TOGGLE_LOCK_DIALOG_EVENT';
export const PEM_GROUP = 'cruduser';
const headingTxt = t`User`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    external_id: t`External ID`,
    sub: t`Sub`,
    email: t`Email`,
    mobile: t`Mobile`,
    first_name: t`First name`,
    last_name: t`Last name`,
    avatar: t`Avatar`,
    admin: t`Admin`,
    locked: t`Locked`,
    locked_reason: t`Locked reason`,
    roles: t`Roles`
});
