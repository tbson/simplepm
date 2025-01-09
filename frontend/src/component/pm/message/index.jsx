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
    const [feature, setFeature] = useState({});
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
                        title: resp.data.project.title
                    },
                    feature: {
                        title: resp.data.title
                    }
                };
                setMessageBreadcrumb(data);
                setFeature({
                    id: resp.data.id,
                    title: resp.data.title,
                    description: resp.data.description
                });
            })
            .catch(RequestUtil.displayError(notification));
    };

    const handleNav = (featureTitle) => {
        const data = { ...messageBreadcrumb, feature: { title: featureTitle } };
        setMessageBreadcrumb(data);
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
                <Chat defaultFeature={feature} onNav={handleNav} />
            ) : null}
        </>
    );
}

Message.displayName = 'Message';
