import * as React from 'react';
import { LoadingOutlined } from '@ant-design/icons';
import { Spin } from 'antd';

export default function Waiting() {
    return (
        <div className="backdrop">
            <Spin indicator={<LoadingOutlined style={{ fontSize: 48 }} spin />} />
        </div>
    );
}
