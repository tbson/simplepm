import React, { useRef, useCallback } from 'react';
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
    const editorCore = useRef(null);

    const handleInitialize = useCallback((instance) => {
        editorCore.current = instance;
    }, []);

    const handleSave = useCallback(() => {
        return editorCore.current.save();
    }, []);
    return (
        <div className={disabled ? 'richtext-view' : 'richtext-edit'}>
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
                minHeight={200}
            />
        </div>
    );
}
