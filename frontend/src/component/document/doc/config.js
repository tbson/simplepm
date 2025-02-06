import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'document/doc',
        endpoints: {
            crud: '',
        }
    }
};

export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_DOC_DIALOG';
export const PEM_GROUP = 'doc';
const headingTxt = t`Doc`;
const name = headingTxt.toLowerCase();
export const getDocs = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    title: t`Title`,
    description: t`Description`,
    task: t`Task`
});
