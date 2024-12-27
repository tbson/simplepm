import * as React from 'react';
import { useState, useEffect } from 'react';
import { useNavigate, useLocation, Outlet, NavLink } from 'react-router-dom';
import { t } from 'ttag';
import { Layout, Menu } from 'antd';
import {
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    UserOutlined,
    TeamOutlined,
    ProjectOutlined,
} from '@ant-design/icons';
import { LOGO_TEXT, DOMAIN } from 'src/const';
import PemUtil from 'service/helper/pem_util';
import NavUtil from 'service/helper/nav_util';
import UserMenu from './user_menu';
import styles from './styles.module.css';

const { Header, Footer, Sider, Content } = Layout;

/**
 * UserLayout.
 */
export default function UserLayout() {
    const navigate = useNavigate();
    const location = useLocation();
    const [menuItems, setMenuItems] = useState([]);

    const [collapsed, setCollapsed] = useState(true);
    const toggle = () => {
        setCollapsed(!collapsed);
    };

    useEffect(() => {
        setMenuItems(getMenuItems());
    }, []);

    const navigateTo = NavUtil.navigateTo(navigate);

    const getMenuItems = () => {
        const result = [];

        PemUtil.canView('crudrole') &&
            result.push({
                label: t`Role`,
                key: `/account/role`,
                icon: <TeamOutlined />
            });
        PemUtil.canView('cruduser') &&
            result.push({
                label: t`User`,
                key: `/account/user`,
                icon: <UserOutlined />
            });
        /*
        PemUtil.canView('crudworkspace') &&
            result.push({
                label: t`Workspace`,
                key: `/pm/workspace`,
                icon: <AppstoreOutlined />
            });
        */
        PemUtil.canView('crudproject') &&
            result.push({
                label: t`Project`,
                key: `/pm/project`,
                icon: <ProjectOutlined />
            });
        return result;
    };

    return (
        <Layout hasSider className={styles.wrapperContainer}>
            <Sider
                trigger={null}
                breakpoint="lg"
                collapsedWidth="42"
                theme="dark"
                collapsible
                collapsed={collapsed}
            >
                <div className="sider">
                    {collapsed || (
                        <div className="logo">
                            <div className="logo-text">
                                <NavLink to="/">{LOGO_TEXT}</NavLink>
                            </div>
                        </div>
                    )}
                    <Menu
                        selectedKeys={[location.pathname]}
                        theme="dark"
                        mode="inline"
                        items={menuItems}
                        onSelect={({ key }) => {
                            navigateTo(key);
                        }}
                    />
                </div>
            </Sider>
            <Layout className="site-layout">
                <Header className="site-layout-header" style={{ padding: 0 }}>
                    <div style={{ display: 'flex' }}>
                        <div style={{ width: 34, paddingLeft: 2, backgroundColor: "white" }}>
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
                        <div style={{ width: 34, backgroundColor: "white" }} >
                            <UserMenu />
                        </div>
                    </div>
                </Header>
                <Content className="site-layout-content">
                    <Outlet />
                </Content>
                <Footer className="layout-footer">
                    <div className="layout-footer-text">
                        Copyright<sup>©</sup> {DOMAIN} {new Date().getFullYear()}
                    </div>
                </Footer>
            </Layout>
        </Layout>
    );
}
