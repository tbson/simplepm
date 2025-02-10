import * as React from 'react';
import { createReactEditorJS } from 'react-editor-js';
import { getTools } from './tools';

/**
 * RichTextInput.
 *
 * @param {Object} props
 * @param {number[]} props.value
 */
export default function RichTextInput({ value, onChange, taskId, disabled = false }) {
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
            tools={getTools(taskId)}
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
