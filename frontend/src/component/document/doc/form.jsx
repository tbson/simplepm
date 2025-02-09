import * as React from 'react';
import { useRef, useEffect } from 'react';
import { App, Form, Input } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import RichTextInput from 'component/common/form/ant/input/richtext_input';
import { urls, getLabels } from './config';

const formName = 'DocForm';
const emptyRecord = {
    id: 0,
    title: '',
    content: '',
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * DocForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function DocForm({ data, onChange }) {
    const { notification } = App.useApp();
    const inputRef = useRef(null);
    const [form] = Form.useForm();

    const labels = getLabels();

    const initialValues = Util.isEmpty(data) ? emptyRecord : data;
    const { id } = initialValues;

    const endPoint = id ? `${urls.crud}${id}` : urls.crud;
    const method = id ? 'put' : 'post';

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
            onFinish={(payload) =>
                FormUtil.submit(endPoint, payload, method)
                    .then((data) => onChange(data, id))
                    .catch(FormUtil.setFormErrors(form, notification))
            }
        >
            <Form.Item name="title" label={labels.title} rules={[FormUtil.ruleRequired()]}>
                <Input ref={inputRef} />
            </Form.Item>

            <Form.Item name="content" label={labels.content}>
                <RichTextInput />
            </Form.Item>
        </Form>
    );
}

DocForm.displayName = formName;
DocForm.formName = formName;
