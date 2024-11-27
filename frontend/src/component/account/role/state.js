import { atom } from 'jotai';
import TableUtil from 'service/helper/table_util';

export const roleOptionSt = atom({
    loaded: false,
    pem: []
});

export const roleTransferSt = atom((get) => {
    const { pem } = get(roleOptionSt);
    return {
        pem: TableUtil.optionToTransfer(pem)
    };
});
