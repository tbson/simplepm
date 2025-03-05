import * as React from 'react';
import { Checkbox } from 'antd';

/**
 * UserInput.
 *
 * @param {Object} props
 * @param {number[]} props.value
 * @param {function} props.onChange
 */
export default function UserInput({ value, onChange, disabled = false }) {
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
