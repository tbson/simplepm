import * as React from 'react';
import { useRef, useEffect } from 'react';
import { Form, Input } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import CheckInput from 'component/common/form/ant/input/check_input';
import { urls, getLabels } from '../config';

const { TextArea } = Input;

const formName = 'AuthClientForm';
const emptyRecord = {
    id: 0,
    uid: '',
    description: '',
    secret: '',
    partition: '',
    default: false
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * AuthClientForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function AuthClientForm({ data, onChange }) {
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
                    .catch(FormUtil.setFormErrors(form))
            }
        >
            <Form.Item name="uid" label={labels.uid} rules={[FormUtil.ruleRequired()]}>
                <Input ref={inputRef} />
            </Form.Item>

            <Form.Item name="description" label={labels.description}>
                <TextArea />
            </Form.Item>

            <Form.Item name="secret" label={labels.secret}>
                <Input type="password" />
            </Form.Item>

            <Form.Item name="partition" label={labels.partition}>
                <Input />
            </Form.Item>

            <Form.Item name="default" label={labels.default}>
                <CheckInput />
            </Form.Item>
        </Form>
    );
}

AuthClientForm.displayName = formName;
AuthClientForm.formName = formName;
