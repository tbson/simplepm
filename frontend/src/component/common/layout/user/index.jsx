import * as React from 'react';
import { useState, useEffect } from 'react';
import { useNavigate, useLocation, Outlet, NavLink } from 'react-router-dom';
import { t } from 'ttag';
import { Layout, Menu, Row, Col, Flex } from 'antd';
import {
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    UserOutlined,
    TagsOutlined,
    TeamOutlined
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

    const [collapsed, setCollapsed] = useState(false);
    const toggle = () => {
        setCollapsed(!collapsed);
    };

    useEffect(() => {
        setMenuItems(getMenuItems());
    }, []);

    const navigateTo = NavUtil.navigateTo(navigate);

    const getMenuItems = () => {
        const result = [];

        result.push({ label: t`Profile`, key: '/', icon: <UserOutlined /> });
        PemUtil.canView('crudrole') &&
            result.push({
                label: t`Role`,
                key: `/account/role`,
                icon: <TagsOutlined />
            });
        PemUtil.canView('cruduser') &&
            result.push({
                label: t`User`,
                key: `/account/user`,
                icon: <TeamOutlined />
            });
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
                                <NavLink to="/">{LOGO_TEXT}</NavLink>
                            </div>
                        </div>
                    )}
                    <Menu
                        className="sidebar-nav"
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
                                <UserMenu />
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
