import { atom } from 'jotai';
import TableUtil from 'service/helper/table_util';

export const taskOptionSt = atom({
    loaded: false,
    feature: []
});

export const taskFilterSt = atom((get) => {
    const { feature } = get(taskOptionSt);
    return {
        feature: TableUtil.optionToFilter(feature)
    };
});
