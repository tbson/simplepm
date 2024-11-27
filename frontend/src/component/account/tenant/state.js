import { atom } from 'jotai';
import TableUtil from 'service/helper/table_util';

export const tenantOptionSt = atom({
    loaded: false,
    auth_client: []
});

export const tenantFilterSt = atom((get) => {
    const { auth_client } = get(tenantOptionSt);
    return {
        auth_client: TableUtil.optionToFilter(auth_client)
    };
});

export const tenantDictSt = atom((get) => {
    const { auth_client } = get(tenantOptionSt);
    const tenantDict = auth_client.reduce((acc, item) => {
        acc[item.value] = item.label;
        return acc;
    }, {});
    return {
        auth_client: tenantDict
    };
});
