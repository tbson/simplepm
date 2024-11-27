import * as React from 'react';
import { useRef, useEffect } from 'react';
import { useAtomValue } from 'jotai';
import { Form, Input } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import SelectInput from 'component/common/form/ant/input/select_input';
import { variableOptionSt } from 'component/config/variable/state';
import { urls, getLabels } from '../config';

const { TextArea } = Input;

const formName = 'VariableForm';
const emptyRecord = {
    id: 0,
    uid: '',
    value: '',
    description: '',
    data_type: 'STRING'
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
    const inputRef = useRef(null);
    const [form] = Form.useForm();
    const variableOption = useAtomValue(variableOptionSt);

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
            <Form.Item name="key" label={labels.key} rules={[FormUtil.ruleRequired()]}>
                <Input ref={inputRef} />
            </Form.Item>

            <Form.Item name="value" label={labels.value}>
                <Input />
            </Form.Item>

            <Form.Item name="description" label={labels.description}>
                <TextArea />
            </Form.Item>

            <Form.Item
                name="data_type"
                label={labels.data_type}
                rules={[FormUtil.ruleRequired()]}
            >
                <SelectInput block options={variableOption.data_type} />
            </Form.Item>
        </Form>
    );
}

VariableForm.displayName = formName;
VariableForm.formName = formName;
