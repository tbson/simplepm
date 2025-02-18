import * as React from 'react';
import { useState, useEffect } from 'react';
import { useNavigate, useLocation, Outlet, NavLink } from 'react-router';
import { t } from 'ttag';
import { Layout, Menu, Badge } from 'antd';
import Icon, {
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    UserOutlined,
    TeamOutlined,
    ProjectOutlined,
    MessageOutlined
} from '@ant-design/icons';
import { LOGO_TEXT, DOMAIN } from 'src/const';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import PemUtil from 'service/helper/pem_util';
import NavUtil from 'service/helper/nav_util';
import UserMenu from './user_menu';
import styles from './styles.module.css';

const { Header, Footer, Sider, Content } = Layout;

function CustomIcon(imageUrl) {
    const url = imageUrl || '/image/default.png';
    return () => <img src={url} alt="icon" style={{ width: 16, height: 16 }} />;
}

/**
 * UserLayout.
 */
export default function UserLayout() {
    const navigate = useNavigate();
    const location = useLocation();
    const [menuItems, setMenuItems] = useState([]);
    const [bookmarkItems, setBookmarkItems] = useState([]);

    const [collapsed, setCollapsed] = useState(true);
    const toggle = () => {
        setCollapsed(!collapsed);
    };

    useEffect(() => {
        setMenuItems(getMenuItems());
        getBookmarks();
    }, [collapsed]);

    useEffect(() => {
        Util.event.listen("FETCH_BOOKMARK", handleFetchBookmark);
        return () => {
            Util.event.remove("FETCH_BOOKMARK", handleFetchBookmark);
        };
    }, []);

    const handleFetchBookmark = () => {
        getBookmarks();
    }

    const navigateTo = NavUtil.navigateTo(navigate);

    const getBookmarks = () => {
        const url = '/pm/project/bookmark/';
        return RequestUtil.apiCall(url)
            .then((resp) => {
                setBookmarkItems(
                    resp.data.map((item) => {
                        return {
                            label: item.title,
                            key: `/pm/project/${item.id}`,
                            icon: (
                                <Badge
                                    size="small"
                                    count={0}
                                    offset={[0, collapsed ? 6 : 0]}
                                >
                                    <Icon component={CustomIcon(item.avatar)} />
                                </Badge>
                            )
                        };
                    })
                );
            })
            .catch(() => {
                setBookmarkItems([]);
            });
    };

    const getMenuItems = () => {
        const result = [];

        PemUtil.canView('user') &&
            result.push({
                label: t`User`,
                key: `/account/user`,
                icon: <UserOutlined />
            });
        PemUtil.canView('role') &&
            result.push({
                label: t`Role`,
                key: `/account/role`,
                icon: <TeamOutlined />
            });
        /*
        PemUtil.canView('workspace') &&
            result.push({
                label: t`Workspace`,
                key: `/pm/workspace`,
                icon: <AppstoreOutlined />
            });
        */
        PemUtil.canView('project') &&
            result.push({
                label: t`Project`,
                key: `/pm/project`,
                icon: <ProjectOutlined />
            });
        /*
        PemUtil.canView('project') &&
            result.push({
                label: t`Message`,
                key: `/event/message`,
                icon: <MessageOutlined />
            });
        */
        return result;
    };

    return (
        <Layout hasSider className={styles.wrapperContainer}>
            {bookmarkItems.length ? (
                <Sider
                    trigger={null}
                    breakpoint="lg"
                    collapsedWidth="42"
                    theme="light"
                    collapsible
                    collapsed={collapsed}
                    className="bookmark-sider"
                >
                    <div className="sider">
                        <Menu
                            selectedKeys={[location.pathname]}
                            theme="light"
                            mode="inline"
                            items={bookmarkItems}
                            onSelect={({ key }) => {
                                navigateTo(key);
                            }}
                        />
                    </div>
                </Sider>
            ) : null}
            <Layout className="site-layout">
                <Header className="site-layout-header" style={{ padding: 0 }}>
                    <div style={{ display: 'flex' }}>
                        <div
                            style={{
                                width: 34,
                                paddingLeft: 2,
                                backgroundColor: 'white'
                            }}
                        >
                            {React.createElement(
                                collapsed ? MenuUnfoldOutlined : MenuFoldOutlined,
                                {
                                    className: 'trigger',
                                    onClick: toggle
                                }
                            )}
                        </div>
                        <div style={{ flex: 1 }}>
                            <Menu
                                mode="horizontal"
                                selectedKeys={[location.pathname]}
                                items={menuItems}
                                onSelect={({ key }) => {
                                    navigateTo(key);
                                }}
                            />
                        </div>
                        <div style={{ width: 34, backgroundColor: 'white' }}>
                            <UserMenu />
                        </div>
                    </div>
                </Header>
                <Content className="site-layout-content">
                    <Outlet />
                </Content>
                {/*
                <Footer className="layout-footer">
                    <div className="layout-footer-text">
                        Copyright<sup>©</sup> {DOMAIN} {new Date().getFullYear()}
                    </div>
                </Footer>
                */}
            </Layout>
        </Layout>
    );
}
