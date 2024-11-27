import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'account/auth-client',
        endpoints: {
            crud: '',
        }
    }
};
export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_AUTH_CLIENT_DIALOG';
export const PEM_GROUP = 'crudauthclient';
const headingTxt = t`Authentication client`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    uid: t`UID`,
    description: t`Description`,
    secret: t`Secret`,
    partition: t`Partition`,
    default: t`Default`
});
