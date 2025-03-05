import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'document/doc',
        endpoints: {
            crud: '',
            createDocFromLink: 'create-doc-from-link'
        }
    },
    task: {
        prefix: 'pm/task',
        endpoints: {
            crud: ''
        }
    }
};

export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const taskUrls = RequestUtil.prefixMapValues(urlMap.task);
export const TOGGLE_LINK_DIALOG_EVENT = 'TOGGLE_DOC_LINK_DIALOG';
export const PEM_GROUP = 'doc';
const headingTxt = t`Doc`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    title: t`Title`,
    description: t`Description`,
    content: t`Content`,
    link: t`Link`
});

export const MODE = {
    EDIT: 'EDIT',
    VIEW: 'VIEW'
};
