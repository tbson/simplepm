import * as React from 'react';
import { createReactEditorJS } from 'react-editor-js';
import { EDITOR_JS_TOOLS } from './tools';

/**
 * RichTextInput.
 *
 * @param {Object} props
 * @param {number[]} props.value
 */
export default function RichTextInput({ value, onChange, disabled = false }) {
    const ReactEditorJS = createReactEditorJS();
    const editorCore = React.useRef(null);

    const handleInitialize = React.useCallback((instance) => {
        editorCore.current = instance;
    }, []);

    const handleSave = React.useCallback(() => {
        return editorCore.current.save();
    }, []);
    return (
        <ReactEditorJS
            readOnly={disabled}
            onInitialize={handleInitialize}
            tools={EDITOR_JS_TOOLS}
            onChange={() => {
                handleSave().then((data) => {
                    onChange(data);
                });
            }}
            defaultValue={value}
            placeholder="Content..."
        />
    );
}
