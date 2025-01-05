import React from 'react';
import { generate, green, presetPalettes, red } from '@ant-design/colors';
import { ColorPicker, theme } from 'antd';
const genPresets = (presets = presetPalettes) =>
    Object.entries(presets).map(([label, colors]) => ({
        label,
        colors
    }));

/**
 * TreeCheckInput.
 *
 * @param {Object} props
 * @param {number[]} props.value
 * @param {function} props.onChange
 */
export default function ColorInput({ value, onChange, disabled = false }) {
    const { token } = theme.useToken();
    const presets = genPresets({
        primary: generate(token.colorPrimary),
        red,
        green
    });
    return (
        <ColorPicker
            defaultFormat="hex"
            disabled={disabled}
            presets={presets}
            value={value}
            onChange={(color) => {
                onChange(color.toHexString());
            }}
        />
    );
}
