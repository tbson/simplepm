import * as React from 'react';
import { useAtomValue } from 'jotai';
import { Form, Input } from 'antd';
import { useParams } from 'react-router-dom';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import CheckInput from 'component/common/form/ant/input/check_input';
import SelectInput from 'component/common/form/ant/input/select_input';
import { userOptionSt } from 'component/account/user/state';
import { urls, getLabels } from '../config';

const { TextArea } = Input;

const formName = 'UserForm';
const emptyRecord = {
    id: 0,
    email: '',
    mobile: '',
    first_name: '',
    last_name: '',
    locked: false,
    locked_reason: '',
    roles: []
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
    const { tenant_id } = useParams();
    const userOption = useAtomValue(userOptionSt);
    const [form] = Form.useForm();

    const labels = getLabels();

    const initialValues = Util.isEmpty(data) ? emptyRecord : data;
    const { id } = initialValues;

    let endPoint = id ? `${urls.crud}${id}` : urls.crud;
    if (tenant_id) {
        endPoint += `?tenant_id=${tenant_id}`;
    }
    const method = id ? 'put' : 'post';

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
            <Form.Item name="email" label={labels.email}>
                <Input disabled />
            </Form.Item>

            <Form.Item name="mobile" label={labels.mobile}>
                <Input disabled />
            </Form.Item>

            <Form.Item name="first_name" label={labels.first_name}>
                <Input disabled />
            </Form.Item>

            <Form.Item name="last_name" label={labels.last_name}>
                <Input disabled />
            </Form.Item>

            <Form.Item name="locked" label={labels.locked}>
                <CheckInput disabled />
            </Form.Item>

            <Form.Item name="locked_reason" label={labels.locked_reason}>
                <TextArea disabled />
            </Form.Item>

            <Form.Item
                name="role_ids"
                label={labels.roles}
                rules={[FormUtil.ruleRequired()]}
            >
                <SelectInput block options={userOption.role} mode="multiple" />
            </Form.Item>
        </Form>
    );
}

UserForm.displayName = formName;
UserForm.formName = formName;
