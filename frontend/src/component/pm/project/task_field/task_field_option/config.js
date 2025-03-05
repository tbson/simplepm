import { t } from 'ttag';

export const TOGGLE_DIALOG_EVENT = 'TOGGLE_TASK_FIELD_OPTION_DIALOG';
const headingTxt = t`Option`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    title: t`Title`,
    description: t`Description`,
    color: t`Color`
});
