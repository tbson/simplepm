import { atom } from 'jotai';
import TableUtil from 'service/helper/table_util';

export const workspaceOptionSt = atom({
    loaded: false,
    auth_client: []
});

export const workspaceFilterSt = atom((get) => {
    const { auth_client } = get(workspaceOptionSt);
    return {
        auth_client: TableUtil.optionToFilter(auth_client)
    };
});

export const workspaceDictSt = atom((get) => {
    const { auth_client } = get(workspaceOptionSt);
    const workspaceDict = auth_client.reduce((acc, item) => {
        acc[item.value] = item.label;
        return acc;
    }, {});
    return {
        auth_client: workspaceDict
    };
});
