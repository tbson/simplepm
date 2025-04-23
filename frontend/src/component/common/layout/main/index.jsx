import React from 'react';
import StorageUtil from 'service/helper/storage_util';
const AdminLayout = (await import('component/common/layout/admin')).default;
const UserLayout = (await import('component/common/layout/user')).default;


export default function MainLayout() {
    const userInfor = StorageUtil.getUserInfo();
    const MainLayout = userInfor?.admin ? AdminLayout : UserLayout;

    return <MainLayout />;
}
