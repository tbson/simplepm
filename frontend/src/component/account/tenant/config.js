import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'account/tenant',
        endpoints: {
            crud: '',
            option: 'option'
        }
    },
    github: {
        prefix: 'git/github',
        endpoints: {
            installUrl: 'install-url'
        }
    }
};
export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const githubUrls = RequestUtil.prefixMapValues(urlMap.github);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_TENANT_DIALOG';
export const PEM_GROUP = 'tenant';
const headingTxt = t`Tenant`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    auth_client_id: t`Auth client`,
    uid: t`UID`,
    title: t`Title`,
    avatar: t`Avatar`
});
