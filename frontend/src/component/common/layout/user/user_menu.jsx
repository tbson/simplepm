import React, { useState } from 'react';
import { t } from 'ttag';
import { UserOutlined, LogoutOutlined } from '@ant-design/icons';
import { Drawer, Avatar, Menu, Divider } from 'antd';
import { useLocation, useNavigate } from 'react-router-dom';
import Util from 'service/helper/util';
import StorageUtil from 'service/helper/storage_util';
import NavUtil from 'service/helper/nav_util';

function MenuHeading({ email, fullName, avatar }) {
    return (
        <div style={{ display: 'flex' }}>
            <div style={{ width: 40 }}>
                <Avatar src={avatar} icon={<UserOutlined />} />
            </div>
            <div style={{ flex: 1 }}>
                <div>
                    <strong>{fullName}</strong>
                </div>
                <div>{email}</div>
            </div>
        </div>
    );
}

export default function UserMenu() {
    const [open, setOpen] = useState(false);
    const navigate = useNavigate();
    const location = useLocation();
    const logout = NavUtil.logout(navigate);
    const navigateTo = NavUtil.navigateTo(navigate);

    const userInfo = StorageUtil.getUserInfo();
    if (!userInfo) {
        return null;
    }
    const { email, avatar, first_name, last_name } = userInfo;
    const fullName = `${first_name} ${last_name}`;

    const handleLogout = () => {
        Util.toggleGlobalLoading();
        logout();
    };

    const handleShow = () => {
        setOpen(true);
    };
    const handleClose = () => {
        setOpen(false);
    };

    return (
        <>
            <Avatar
                src={avatar}
                icon={<UserOutlined />}
                onClick={handleShow}
                className="pointer"
            />
            <Drawer
                closeIcon={null}
                title={null}
                onClose={handleClose}
                open={open}
            >
                <MenuHeading email={email} fullName={fullName} avatar={avatar} />
                <Divider />
                <Menu
                    inlineIndent={6}
                    selectedKeys={[location.pathname]}
                    mode="inline"
                    items={[
                        { label: t`Your profile`, key: '/', icon: <UserOutlined /> },
                        {
                            label: t`Sign out`,
                            key: `/logout`,
                            icon: <LogoutOutlined />
                        }
                    ]}
                    onSelect={({ key }) => {
                        if (key === '/logout') {
                            handleLogout();
                            return;
                        }
                        navigateTo(key);
                        handleClose();
                    }}
                />
            </Drawer>
        </>
    );
}
