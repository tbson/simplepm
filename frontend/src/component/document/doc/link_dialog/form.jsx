import React, { useRef, useEffect } from 'react';
import { App, Form, Input } from 'antd';
import { useParams } from 'react-router';
import FormUtil from 'service/helper/form_util';
import { urls, getLabels } from '../config';

const formName = 'DocLinkForm';
const emptyRecord = {
    link: ''
};

/**
 * DocLinkForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function DocLinkForm({ onChange }) {
    const taskId = parseInt(useParams().taskId, 10);
    const { notification } = App.useApp();
    const inputRef = useRef(null);
    const [form] = Form.useForm();

    const labels = getLabels();

    const initialValues = emptyRecord;

    useEffect(() => {
        inputRef.current.focus({ cursor: 'end' });
    }, []);

    return (
        <Form
            form={form}
            name={formName}
            colon={false}
            labelWrap
            labelCol={{ span: 6 }}
            wrapperCol={{ span: 18 }}
            initialValues={{ ...initialValues }}
            onFinish={(data) => {
                const payload = {
                    ...data,
                    task_id: taskId
                };
                FormUtil.submit(urls.createDocFromLink, payload, 'post')
                    .then((data) => onChange(data))
                    .catch(FormUtil.setFormErrors(form, notification));
            }}
        >
            <Form.Item
                name="link"
                label={labels.link}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input ref={inputRef} />
            </Form.Item>
        </Form>
    );
}

DocLinkForm.displayName = formName;
DocLinkForm.formName = formName;
