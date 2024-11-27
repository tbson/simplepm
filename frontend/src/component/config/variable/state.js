import { atom } from 'jotai';
import TableUtil from 'service/helper/table_util';

export const variableOptionSt = atom({
    loaded: false,
    data_type: []
});

export const variableFilterSt = atom((get) => {
    const { data_type } = get(variableOptionSt);
    return {
        data_type: TableUtil.optionToFilter(data_type)
    };
});
