import React, { useEffect } from 'react';
import { useParams } from 'react-router';
import { useAtom } from 'jotai';
import { Skeleton } from 'antd';
import RequestUtil from 'service/helper/request_util';
import { taskOptionSt } from './state';
import { urls } from './config';
import TaskKanban from './kanban';

export default function Task() {
    const projectId = parseInt(useParams().projectId, 10);
    const [taskOption, setTaskOption] = useAtom(taskOptionSt);
    useEffect(() => {
        setTaskOption({ ...taskOption, loaded: false });
        getOption(projectId);
    }, [projectId]);

    const getOption = (projectId) => {
        RequestUtil.apiCall(urls.option, { project_id: projectId })
            .then((resp) => {
                setTaskOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setTaskOption((prev) => ({ ...prev, loaded: true }));
            });
    };
    if (!taskOption.loaded) {
        return <Skeleton loading={true} active />;
    }
    return (
        <div key={projectId}>
            <TaskKanban
                project={{ id: projectId, title: taskOption.project_info.title }}
            />
        </div>
    );
}

Task.displayName = 'Task';
