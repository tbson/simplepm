import React, { useRef, useEffect } from 'react';
import { t } from 'ttag';
import { App, Form, Input } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import { urls } from '../config';

const formName = 'VariableForm';
const emptyRecord = {
    email: '',
    password: '',
    password_again: '',
    code: ''
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * VariableForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function VariableForm({ data, onChange }) {
    const { notification } = App.useApp();
    const inputRef = useRef(null);
    const [form] = Form.useForm();

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
            labelCol={{ span: 8 }}
            wrapperCol={{ span: 16 }}
            initialValues={{ ...initialValues }}
            onFinish={(payload) =>
                FormUtil.submit(endPoint, payload, method)
                    .then((data) => onChange(data, id))
                    .catch(FormUtil.setFormErrors(form, notification))
            }
        >
            <Form.Item name="email" label={t`Email`} rules={[FormUtil.ruleRequired()]}>
                <Input ref={inputRef} />
            </Form.Item>

            <Form.Item
                name="password"
                label={t`Password`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input type="password" />
            </Form.Item>

            <Form.Item
                name="password_again"
                label={t`Password again`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input type="password_again" />
            </Form.Item>

            <Form.Item name="code" label={t`Code`} rules={[FormUtil.ruleRequired()]}>
                <Input type="code" />
            </Form.Item>
        </Form>
    );
}

VariableForm.displayName = formName;
VariableForm.formName = formName;
