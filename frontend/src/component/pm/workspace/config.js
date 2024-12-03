import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'pm/workspace',
        endpoints: {
            crud: '',
        }
    }
};
export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_WORKSPACE_DIALOG';
export const PEM_GROUP = 'crudworkspace';
const headingTxt = t`Workspace`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    tenant_id: t`Tenant`,
    title: t`Title`,
    description: t`Description`,
    avatar: t`Avatar`,
    order: t`Order`,
});
