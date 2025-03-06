import React from 'react';
import { t } from 'ttag';
import { Button, Tooltip } from 'antd';
import {
    EditOutlined,
    EyeOutlined,
    DeleteOutlined,
    GlobalOutlined,
    CheckOutlined,
    PlusOutlined
} from '@ant-design/icons';

export function IconButton({
    icon,
    tootip,
    onClick,
    value,
    type = 'default',
    disabled = false
}) {
    return (
        <Tooltip title={tootip}>
            <Button
                type={type}
                disabled={disabled}
                htmlType="button"
                icon={icon}
                size="small"
                onClick={onClick.bind(null, value)}
            />
        </Tooltip>
    );
}

export function AddNewBtn({ onClick, value }) {
    return (
        <Button
            type="primary"
            icon={<PlusOutlined />}
            onClick={onClick.bind(null, value)}
        >
            {t`Add new`}
        </Button>
    );
}

export function RemoveSelectedBtn({ onClick, value=[] }) {
    return (
        <Button
            type="primary"
            danger
            icon={<DeleteOutlined />}
            disabled={!value.length}
            onClick={onClick.bind(null, value)}
        >
            {t`Remove selected`}
        </Button>
    );
}

export function EditBtn({ onClick, value }) {
    return (
        <Tooltip title={t`Update`}>
            <Button
                type="default"
                htmlType="button"
                icon={<EditOutlined />}
                size="small"
                onClick={onClick.bind(null, value)}
            />
        </Tooltip>
    );
}

export function RemoveBtn({ onClick, value }) {
    return (
        <Tooltip title={t`Remove`}>
            <Button
                danger
                type="default"
                htmlType="button"
                icon={<DeleteOutlined />}
                size="small"
                onClick={onClick.bind(null, value)}
            />
        </Tooltip>
    );
}

export function ViewBtn({ onClick, value }) {
    return (
        <Tooltip title={t`View`}>
            <Button
                type="default"
                htmlType="button"
                icon={<EyeOutlined />}
                size="small"
                onClick={onClick.bind(null, value)}
            />
        </Tooltip>
    );
}

export function LinkBtn({ onClick, value }) {
    return (
        <Tooltip title={t`Link`}>
            <Button
                type="default"
                htmlType="button"
                icon={<GlobalOutlined />}
                size="small"
                onClick={onClick.bind(null, value)}
            />
        </Tooltip>
    );
}

export function CheckBtn({ onClick, disabled, value }) {
    return (
        <Tooltip title={t`Check`}>
            <Button
                type="default"
                htmlType="button"
                icon={<CheckOutlined />}
                disabled={disabled}
                size="small"
                onClick={onClick.bind(null, value)}
            />
        </Tooltip>
    );
}
