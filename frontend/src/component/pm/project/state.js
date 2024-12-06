import { atom } from 'jotai';
import TableUtil from 'service/helper/table_util';

export const projectOptionSt = atom({
    loaded: false,
    workspace: [],
    layout: []
});

export const projectFilterSt = atom((get) => {
    const { workspace, layout } = get(projectOptionSt);
    return {
        workspace: TableUtil.optionToFilter(workspace),
        layout: TableUtil.optionToFilter(layout)
    };
});
