import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'pm/task-field',
        endpoints: {
            crud: '',
            reorder: 'reorder',
        }
    }
};
export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const TOGGLE_TASK_FIELD_EVENT = 'TOGGLE_TASK_FIELD';
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_TASK_FIELD_DIALOG';
export const PEM_GROUP = 'crudtaskfield';
const headingTxt = t`Task Field`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    title: t`Title`,
    description: t`Description`,
    type: t`Type`,
    order: t`Order`,
});
