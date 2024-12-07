import * as React from 'react';
import DatePicker from 'component/common/form/ant/date_picker';
import { DATE_REABLE_FORMAT } from 'service/helper/date_util';

/**
 * DateInput.
 *
 * @param {Object} props
 * @param {string} props.value
 * @param {function} props.onChange
 * @param {string} props.label
 * @returns {ReactElement}
 */
export default function DateInput({ value, onChange }) {
    const handleChange = (date) => {
        const now = new Date();
        date.setHours(
            now.getHours(),
            now.getMinutes(),
            now.getSeconds(),
            now.getMilliseconds()
        );
        onChange(date);
    };
    return (
        <DatePicker
            value={value}
            onChange={handleChange}
            format={DATE_REABLE_FORMAT}
            style={{ width: '100%' }}
        />
    );
}
