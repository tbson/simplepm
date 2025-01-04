import * as React from 'react';
import { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { useAtom } from 'jotai';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { taskOptionSt } from './state';
import { urls, getMessages } from './config';
import TaskKanban from './kanban';
import FeatureTable from 'component/pm/feature/table';

export default function Task() {
    const { project_id } = useParams();
    const projectId = parseInt(project_id);
    const [taskOption, setTaskOption] = useAtom(taskOptionSt);
    useEffect(() => {
        if (!taskOption.loaded) {
            getOption();
        }
    }, []);

    const getOption = () => {
        RequestUtil.apiCall(urls.option, { project_id })
            .then((resp) => {
                setTaskOption({ ...resp.data, loaded: true });
            })
            .catch(() => {
                setTaskOption((prev) => ({ ...prev, loaded: true }));
            });
    };

    const messages = getMessages();
    return (
        <>
            <PageHeading>
                <>{messages.heading}</>
            </PageHeading>
            <FeatureTable projectId={projectId} />
            <br />
            <TaskKanban projectId={projectId} />
        </>
    );
}

Task.displayName = 'Task';
