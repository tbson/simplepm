import * as React from 'react';
import { useRef, useEffect } from 'react';
import { useAtomValue } from 'jotai';
import { Form, Input } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import SelectInput from 'component/common/form/ant/input/select_input';
import { projectOptionSt } from 'component/pm/project/state';
import { getLabels } from '../config';

const { TextArea } = Input;

const formName = 'TaskFieldOptionForm';
const emptyRecord = {
    id: 0,
    title: '',
    description: '',
    color: 'GRAY'
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * TaskFieldOptionForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function TaskFieldOptionForm({ data, onChange }) {
    const inputRef = useRef(null);
    const [form] = Form.useForm();
    const projectOption = useAtomValue(projectOptionSt);
    const labels = getLabels();

    const initialValues = Util.isEmpty(data) ? emptyRecord : data;
    const { id } = initialValues;

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
            onFinish={(payload) => onChange(payload, id)}
        >
            <Form.Item
                name="title"
                label={labels.title}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input ref={inputRef} />
            </Form.Item>

            <Form.Item name="description" label={labels.description}>
                <TextArea />
            </Form.Item>

            <Form.Item
                name="color"
                label={labels.color}
                rules={[FormUtil.ruleRequired()]}
            >
                <SelectInput block options={projectOption.task_field.color} />
            </Form.Item>
        </Form>
    );
}

TaskFieldOptionForm.displayName = formName;
TaskFieldOptionForm.formName = formName;
