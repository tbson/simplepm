import React, { useEffect, useState } from 'react';
import { Link, useParams, useNavigate } from 'react-router';
import { createStyles } from 'antd-style';
import { App, Breadcrumb, Skeleton } from 'antd';
import PageHeading from 'component/common/page_heading';
import RequestUtil from 'service/helper/request_util';
import NavUtil from 'service/helper/nav_util';
import DocTable from './table';
import DocForm from './form';
import DocView from './view';
import { getStyles } from './style';
import { urls, taskUrls, MODE } from './config';

export default function Doc() {
    const { notification } = App.useApp();
    const useStyle = getStyles(createStyles);
    const { styles } = useStyle();
    const { taskId, docId } = useParams();
    const [project, setProject] = useState({});
    const [task, setTask] = useState({});
    const [init, setInit] = useState(true);
    const [mode, setMode] = useState(docId ? MODE.VIEW : MODE.EDIT);
    const [doc, setDoc] = useState({});
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

    const navigateTo = NavUtil.navigateTo(useNavigate());

    useEffect(() => {
        getBreadcrumb();
    }, []);

    useEffect(() => {
        if (!docId) {
            setInit(false);
            return;
        }
        getDetail();
    }, [docId]);

    const getDetail = () => {
        RequestUtil.apiCall(`${urls.crud}${docId}`)
            .then((resp) => {
                setDoc(resp.data);
            })
            .catch(RequestUtil.displayError(notification))
            .finally(() => {
                setInit(false);
            });
    };

    const getBreadcrumb = () => {
        RequestUtil.apiCall(`${taskUrls.crud}${taskId}`)
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

    const renderDisplay = () => {
        if (init) {
            return <Skeleton active />;
        }
        if (mode === MODE.EDIT) {
            return (
                <DocForm
                    data={doc}
                    onChange={() => {
                        navigateTo(`/pm/task/${taskId}`);
                    }}
                />
            );
        }
        return (
            <DocView
                data={doc}
                toggleMode={() => {
                    setMode(MODE.EDIT);
                }}
            />
        );
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
                                <Link to={`/pm/task/${taskId}`}>
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
                    <div className="flex-item-remaining">{renderDisplay()}</div>
                </div>
            ) : null}
        </div>
    );
}

Doc.displayName = 'Doc';
