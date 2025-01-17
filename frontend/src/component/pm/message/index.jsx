import * as React from 'react';
import { useEffect, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import { App, Breadcrumb, Skeleton } from 'antd';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { taskUrls } from './config';
import Chat from './chat';

export default function Message() {
    const { notification } = App.useApp();
    const { project_id, task_id } = useParams();
    const [task, setTask] = useState({});
    const [messageBreadcrumb, setMessageBreadcrumb] = useState({
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
                const data = {
                    loaded: true,
                    project: {
                        title: resp.data.project.title
                    },
                    task: {
                        title: resp.data.title
                    }
                };
                setMessageBreadcrumb(data);
                setTask({
                    id: resp.data.id,
                    title: resp.data.title,
                    description: resp.data.description
                });
            })
            .catch(RequestUtil.displayError(notification));
    };

    const handleNav = (taskTitle) => {
        const data = { ...messageBreadcrumb, task: { title: taskTitle } };
        setMessageBreadcrumb(data);
    };

    return (
        <div className="flex-column flex-item-remaining">
            <PageHeading>
                <Breadcrumb
                    items={[
                        {
                            title: (
                                <Link to={`/pm/task/${project_id}`}>
                                    {messageBreadcrumb.project.title}
                                </Link>
                            )
                        },
                        {
                            title: messageBreadcrumb.task.title
                        }
                    ]}
                />
            </PageHeading>
            <Skeleton loading={!messageBreadcrumb.loaded} active />
            {messageBreadcrumb.loaded ? (
                <Chat defaultTask={task} onNav={handleNav} />
            ) : null}
        </div>
    );
}

Message.displayName = 'Message';
