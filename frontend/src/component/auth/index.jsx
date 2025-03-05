import React from 'react';
import { Outlet } from 'react-router';

export default function Auth() {
    return <Outlet />;
}

Auth.displayName = 'Auth';
