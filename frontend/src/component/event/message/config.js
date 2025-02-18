import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'event/message',
        endpoints: {
            crud: '',
            delete: '/delete',
        }
    },
    task: {
        prefix: 'pm/task',
        endpoints: {
            option: 'option',
            crud: '',
        }
    },
};

export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const taskUrls = RequestUtil.prefixMapValues(urlMap.task);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_MESSAGE_DIALOG';
export const PEM_GROUP = 'message';
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
