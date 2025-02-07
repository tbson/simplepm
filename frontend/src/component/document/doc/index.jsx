import * as React from 'react';
import { useEffect, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import { createStyles } from 'antd-style';
import { App, Breadcrumb, Skeleton } from 'antd';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { taskUrls } from './config';
import DocTable from './table';
import { getStyles } from './style';

export default function Doc() {
    const { notification } = App.useApp();
    const useStyle = getStyles(createStyles);
    const { styles } = useStyle();
    const { task_id } = useParams();
    const [project, setProject] = useState({});
    const [task, setTask] = useState({});
    const [breadcrumb, setBreadcrumb] = useState({
        project: {
            id: null,
            title: 'Loading...',
            description: ''
        },
        task: {
            id: null,
            title: 'Loading...',
            description: ''
        }
    });
    useEffect(() => {
        getBreadcrumb();
    }, []);

    const getBreadcrumb = () => {
        RequestUtil.apiCall(`${taskUrls.crud}${task_id}`)
            .then((resp) => {
                setProject(resp.data.project);
                const data = {
                    loaded: true,
                    project: {
                        title: resp.data.project.title
                    },
                    task: {
                        title: resp.data.title
                    }
                };
                setBreadcrumb(data);
                setTask({
                    id: resp.data.id,
                    title: resp.data.title,
                    description: resp.data.description
                });
            })
            .catch(RequestUtil.displayError(notification));
    };

    const handleNav = (taskTitle) => {
        const data = { ...breadcrumb, task: { title: taskTitle } };
        setBreadcrumb(data);
    };

    return (
        <div className="flex-column flex-item-remaining">
            <PageHeading>
                <Breadcrumb
                    items={[
                        {
                            title: (
                                <Link to={`/pm/project/${project.id}`}>
                                    {breadcrumb.project.title}
                                </Link>
                            )
                        },
                        {
                            title: (
                                <Link to={`/pm/task/${task_id}`}>
                                    {breadcrumb.task.title}
                                </Link>
                            )
                        },
                        {
                            title: 'New document'
                        }
                    ]}
                />
            </PageHeading>
            <Skeleton loading={!breadcrumb.loaded} active />
            {breadcrumb.loaded ? (
                <div className={styles.layout}>
                    <DocTable taskId={task.id} />
                    <div className="flex-item-remaining">Doc display here......</div>
                </div>
            ) : null}
        </div>
    );
}

Doc.displayName = 'Doc';
