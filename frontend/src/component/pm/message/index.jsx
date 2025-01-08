import * as React from 'react';
import { useEffect, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import { App, Breadcrumb, Skeleton } from 'antd';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import { featureUrls } from './config';
import Chat from './chat';

export default function Message() {
    const { notification } = App.useApp();
    const { project_id, feature_id } = useParams();
    const [messageBreadcrumb, setMessageBreadcrumb] = useState({
        project: {
            id: null,
            title: 'Loading...',
            description: ''
        },
        feature: {
            id: null,
            title: 'Loading...',
            description: ''
        }
    });
    useEffect(() => {
        getBreadcrumb();
    }, []);

    const getBreadcrumb = () => {
        RequestUtil.apiCall(`${featureUrls.crud}${feature_id}`)
            .then((resp) => {
                const data = {
                    loaded: true,
                    project: {
                        id: resp.data.project.id,
                        title: resp.data.project.title,
                        description: resp.data.project.description
                    },
                    feature: {
                        id: resp.data.id,
                        title: resp.data.title,
                        description: resp.data.description
                    }
                };
                setMessageBreadcrumb(data);
            })
            .catch(RequestUtil.displayError(notification));
    };

    return (
        <>
            <PageHeading>
                <Breadcrumb
                    items={[
                        {
                            title: <Link to={`/pm/project`}>Project</Link>
                        },
                        {
                            title: (
                                <Link to={`/pm/task/${project_id}`}>
                                    {messageBreadcrumb.project.title}
                                </Link>
                            )
                        },
                        {
                            title: messageBreadcrumb.feature.title
                        }
                    ]}
                />
            </PageHeading>
            <Skeleton loading={!messageBreadcrumb.loaded} active />
            {messageBreadcrumb.loaded ? (
                <Chat defaultFeature={messageBreadcrumb.feature} />
            ) : null}
        </>
    );
}

Message.displayName = 'Message';
