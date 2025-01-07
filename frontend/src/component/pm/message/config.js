import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'socket/get-jwt',
        endpoints: {
            getJwt: '',
        }
    },
    feature: {
        prefix: 'pm/feature',
        endpoints: {
            crud: '',
        }
    }
};

export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const featureUrls = RequestUtil.prefixMapValues(urlMap.feature);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_MESSAGE_DIALOG';
export const PEM_GROUP = 'crudmessage';
const headingTxt = t`Message`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    title: t`Title`,
    description: t`Description`,
    feature: t`Feature`
});
