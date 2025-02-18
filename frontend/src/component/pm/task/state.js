import { atom } from 'jotai';

export const taskOptionSt = atom({
    loaded: false,
    project_info: {},
    feature: [],
    status: [],
    task_field: [],
    user: [],
});
