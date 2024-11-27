import { atom } from 'jotai';

export const accountOptionSt = atom({
    loaded: false,
    role: [],
    pem: [],
    profile_type: []
});
