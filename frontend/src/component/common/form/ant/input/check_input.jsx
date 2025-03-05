import React from 'react';
import { Checkbox } from 'antd';

/**
 * CheckInput.
 *
 * @param {Object} props
 * @param {number[]} props.value
 * @param {function} props.onChange
 */
export default function CheckInput({ value, onChange, disabled = false }) {
    return (
        <Checkbox
            disabled={disabled}
            checked={value}
            onChange={(e) => {
                onChange(e.target.checked);
            }}
        />
    );
}
