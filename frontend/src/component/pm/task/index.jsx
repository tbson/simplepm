import * as React from 'react';
import { useEffect } from 'react';
import { useParams, Link } from 'react-router';
import { useAtom } from 'jotai';
import { Breadcrumb, Skeleton } from 'antd';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { taskOptionSt } from './state';
import { urls } from './config';
import TaskKanban from './kanban';
import FeatureTable from 'component/pm/feature/table';

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
            <PageHeading>
                <Breadcrumb
                    items={[
                        {
                            title: taskOption.project_info.title
                        }
                    ]}
                />
            </PageHeading>
            {/*
            <FeatureTable projectId={projectId} />
            <br />
            */}
            <TaskKanban projectId={projectId} />
        </div>
    );
}

Task.displayName = 'Task';
