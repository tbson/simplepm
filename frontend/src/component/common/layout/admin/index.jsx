import React, { useState, useEffect } from 'react';
import { useNavigate, useLocation, Outlet, NavLink } from 'react-router';
import { t } from 'ttag';
import { Layout, Menu, Row, Col, Button, Flex } from 'antd';
import {
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    UserOutlined,
    LogoutOutlined,
    SettingFilled,
    AuditOutlined,
} from '@ant-design/icons';
import { LOGO_TEXT, DOMAIN } from 'src/const';
import StorageUtil from 'service/helper/storage_util';
import PemUtil from 'service/helper/pem_util';
import NavUtil from 'service/helper/nav_util';
import LocaleSelect from 'component/common/locale_select.jsx';
import styles from './styles.module.css';

const { Header, Footer, Sider, Content } = Layout;

/**
 * processSelectedKeys.
 *
 * @param {string} pathname
 * @returns {string}
 */
function processSelectedKeys(pathname) {
    if (pathname.startsWith('/user')) {
        return '/user';
    }
    return [pathname];
}

/**
 * AdminLayout.
 */
export default function AdminLayout() {
    const navigate = useNavigate();
    const location = useLocation();
    const [menuItems, setMenuItems] = useState([]);
    const [selectedKeys, setSelectedKeys] = useState(
        processSelectedKeys(location.pathname)
    );

    const [collapsed, setCollapsed] = useState(false);
    const toggle = () => {
        setCollapsed(!collapsed);
    };

    useEffect(() => {
        setMenuItems(getMenuItems());
    }, []);

    const logout = NavUtil.logout(navigate);
    const navigateTo = NavUtil.navigateTo(navigate);

    const getMenuItems = () => {
        const result = [];

        result.push({ label: t`Profile`, key: '/', icon: <UserOutlined /> });
        PemUtil.canView('tenant') &&
            result.push({
                label: t`Tenant`,
                key: '/account/tenant',
                icon: <AuditOutlined />
            });
        PemUtil.canView('variable') &&
            result.push({
                label: t`Variable`,
                key: '/config/variable',
                icon: <SettingFilled />
            });
        /*
        PemUtil.canView('user') &&
            result.push({
                label: t`User`,
                key: '/account/user',
                icon: <TeamOutlined />
            });
        PemUtil.canView('role') &&
            result.push({
                label: t`Role`,
                key: '/account/role',
                icon: <TagsOutlined />
            });
        */
        return result;
    };

    return (
        <Layout hasSider className={styles.wrapperContainer}>
            <Sider
                trigger={null}
                breakpoint="lg"
                collapsedWidth="34"
                theme="dark"
                collapsible
                collapsed={collapsed}
                onBreakpoint={(broken) => {
                    setCollapsed(broken);
                }}
            >
                <div className="sider">
                    {collapsed || (
                        <div className="logo">
                            <div className="logo-text">
                                <NavLink
                                    to="/"
                                    onClick={() => {
                                        setSelectedKeys(['/']);
                                    }}
                                >
                                    {LOGO_TEXT}
                                </NavLink>
                            </div>
                        </div>
                    )}
                    <Menu
                        className="sidebar-nav"
                        selectedKeys={selectedKeys}
                        theme="dark"
                        mode="inline"
                        items={menuItems}
                        onSelect={({ key }) => {
                            navigateTo(key);
                            setSelectedKeys([key]);
                        }}
                    />
                </div>
            </Sider>
            <Layout className="site-layout">
                <Header className="site-layout-header" style={{ padding: 0 }}>
                    <Row>
                        <Col span={12}>
                            {React.createElement(
                                collapsed ? MenuUnfoldOutlined : MenuFoldOutlined,
                                {
                                    className: 'trigger',
                                    onClick: toggle
                                }
                            )}
                        </Col>
                        <Col span={12} style={{ paddingRight: 5 }}>
                            <Flex gap={5} justify="flex-end">
                                <span>{StorageUtil.getUserInfo().first_name}</span>
                                <LocaleSelect />
                                <Button
                                    icon={<LogoutOutlined />}
                                    onClick={() => {
                                        logout();
                                    }}
                                    danger
                                />
                            </Flex>
                        </Col>
                    </Row>
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
