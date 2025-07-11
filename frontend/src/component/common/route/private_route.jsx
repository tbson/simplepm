import React from 'react';
import { Outlet, Navigate } from 'react-router';
import StorageUtil from 'service/helper/storage_util';

export default function PrivateRoute() {
    return StorageUtil.getUserInfo() ? <Outlet /> : <Navigate to="/login" />;
}

PrivateRoute.displayName = 'PrivateRoute';
