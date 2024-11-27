import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'account/role',
        endpoints: {
            crud: '',
            option: 'option'
        }
    }
};
export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_TENANT_DIALOG';
export const PEM_GROUP = 'crudrole';
const headingTxt = t`Role`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    title: t`Title`,
    pem_ids: t`Permissions`,
});
