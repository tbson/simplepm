import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'socket',
        endpoints: {
            publishMessage: 'publish-message',
        }
    },
    task: {
        prefix: 'pm/task',
        endpoints: {
            crud: '',
        }
    },
    message: {
        prefix: 'pm/message',
        endpoints: {
            crud: '',
        }
    }
};

export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const taskUrls = RequestUtil.prefixMapValues(urlMap.task);
export const messageUrls = RequestUtil.prefixMapValues(urlMap.message);
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
    task: t`Task`
});
