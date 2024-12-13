import { atom } from 'jotai';
import TableUtil from 'service/helper/table_util';

export const projectOptionSt = atom({
    loaded: false,
    workspace: [],
    layout: [],
    status: [],
    task_field: {
        type: [],
        color: []
    }
});

export const projectFilterSt = atom((get) => {
    const { workspace, layout, status } = get(projectOptionSt);
    return {
        workspace: TableUtil.optionToFilter(workspace),
        layout: TableUtil.optionToFilter(layout),
        status: TableUtil.optionToFilter(status)
    };
});
