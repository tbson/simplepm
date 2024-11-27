import * as React from 'react';
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

export function AddNewBtn({ onClick }) {
    return (
        <Button type="primary" icon={<PlusOutlined />} onClick={onClick}>
            {t`Add new`}
        </Button>
    );
}

export function RemoveSelectedBtn({ ids, onClick }) {
    return (
        <Button
            type="primary"
            danger
            icon={<DeleteOutlined />}
            disabled={!ids.length}
            onClick={() => onClick(ids)}
        >
            {t`Remove selected`}
        </Button>
    );
}

export function EditBtn({ onClick }) {
    return (
        <Tooltip title={t`Update`}>
            <Button
                type="default"
                htmlType="button"
                icon={<EditOutlined />}
                size="small"
                title="hello"
                onClick={onClick}
            />
        </Tooltip>
    );
}

export function RemoveBtn({ onClick }) {
    return (
        <Tooltip title={t`Remove`}>
            <Button
                danger
                type="default"
                htmlType="button"
                icon={<DeleteOutlined />}
                size="small"
                onClick={onClick}
            />
        </Tooltip>
    );
}

export function ViewBtn({ onClick }) {
    return (
        <Tooltip title={t`View`}>
            <Button
                type="default"
                htmlType="button"
                icon={<EyeOutlined />}
                size="small"
                onClick={onClick}
            />
        </Tooltip>
    );
}

export function LinkBtn({ onClick }) {
    return (
        <Tooltip title={t`Link`}>
            <Button
                type="default"
                htmlType="button"
                icon={<GlobalOutlined />}
                size="small"
                onClick={onClick}
            />
        </Tooltip>
    );
}

export function CheckBtn({ onClick, disabled }) {
    return (
        <Tooltip title={t`Check`}>
            <Button
                type="default"
                htmlType="button"
                icon={<CheckOutlined />}
                disabled={disabled}
                size="small"
                onClick={onClick}
            />
        </Tooltip>
    );
}
