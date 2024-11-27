import * as React from 'react';
import { Form, Input } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import CheckInput from 'component/common/form/ant/input/check_input';
import { lockUrls, getLabels } from '../config';

const { TextArea } = Input;

const formName = 'UserForm';
const emptyRecord = {
    id: 0,
    locked: false,
    locked_reason: ''
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * UserForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function UserForm({ data, onChange }) {
    const [form] = Form.useForm();
    const locked = Form.useWatch('locked', { form, preserve: true });

    const labels = getLabels();

    const initialValues = Util.isEmpty(data) ? emptyRecord : data;
    const { id } = initialValues;

    const endPoint = `${lockUrls.lock}${id}`;
    const method = 'put';

    const handleValuesChange = (changedValues, _allValues) => {
        if ('locked' in changedValues && !changedValues.locked) {
            form.setFieldsValue({ locked_reason: '' });
        }
    };

    return (
        <Form
            form={form}
            name={formName}
            colon={false}
            labelWrap
            labelCol={{ span: 6 }}
            wrapperCol={{ span: 18 }}
            initialValues={{ ...initialValues }}
            onValuesChange={handleValuesChange}
            onFinish={(payload) =>
                FormUtil.submit(endPoint, payload, method)
                    .then((data) => onChange(data, id))
                    .catch(FormUtil.setFormErrors(form))
            }
        >
            <Form.Item name="locked" label={labels.locked}>
                <CheckInput />
            </Form.Item>

            <Form.Item name="locked_reason" label={labels.locked_reason}>
                <TextArea disabled={!locked} />
            </Form.Item>
        </Form>
    );
}

UserForm.displayName = formName;
UserForm.formName = formName;
