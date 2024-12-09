import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'pm/project',
        endpoints: {
            crud: '',
            option: 'option'
        }
    }
};
export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_VARIABLE_DIALOG';
export const PEM_GROUP = 'crudproject';
const headingTxt = t`Project`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    workspace_id: t`Workspace`,
    title: t`Title`,
    description: t`Description`,
    avatar: t`Avatar`,
    layout: t`Layout`,
    order: t`Order`,
    status: t`Status`,
    finished_at: t`Finished at`,
    created_at: t`Created at`,
    updated_at: t`Updated at`
});
