import * as React from 'react';
import { t } from 'ttag';
import { Button, Divider } from 'antd';
import { LeftOutlined, RightOutlined } from '@ant-design/icons';

/**
 * @callback onChange
 * @param {String} url
 * @returns {void}
 */

export const defaultPages = {
    next: '',
    prev: ''
};

/**
 * ActionBtn.
 *
 * @param {Object} props
 * @param {string} props.type
 * @param {number} props.page
 * @param {onChange} props.onChange
 */
function ActionBtn({ type, page, onChange }) {
    if (!page) return null;
    const label = {
        prev: t`Prev`,
        next: t`Next`
    };
    return (
        <Button
            type="primary"
            key={1}
            icon={type === 'prev' ? <LeftOutlined /> : <RightOutlined />}
            onClick={() => onChange(page)}
        >
            {label[type]}
        </Button>
    );
}

/**
 * Pagination.
 *
 * @param {Object} props
 * @param {string} props.next
 * @param {string} props.prev
 * @param {onChange} props.onChange
 * @returns {ReactElement}
 */
export default function Pagination({ next, prev, onChange }) {
    return (
        <div className="right">
            <ActionBtn type="prev" page={prev} onChange={onChange} />
            {(next && prev) ? <Divider type="vertical" /> : null}
            <ActionBtn type="next" page={next} onChange={onChange} />
        </div>
    );
}
