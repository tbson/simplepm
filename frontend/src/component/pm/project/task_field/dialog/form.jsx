import React, { useRef, useEffect } from 'react';
import { useAtomValue } from 'jotai';
import { App, Form, Input } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import SelectInput from 'component/common/form/ant/input/select_input';
import { projectIdSt } from '../state';
import { projectOptionSt } from 'component/pm/project/state';
import TaskFieldOptionTable from '../task_field_option/table';
import { urls, getLabels } from '../config';

const { TextArea } = Input;

const formName = 'TaskFieldForm';
const emptyRecord = {
    id: 0,
    title: '',
    description: '',
    type: 'TEXT',
    task_field_options: []
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * TaskFieldForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function TaskFieldForm({ data, onChange }) {
    const { notification } = App.useApp();
    const inputRef = useRef(null);
    const [form] = Form.useForm();
    const projectId = useAtomValue(projectIdSt);
    const projectOption = useAtomValue(projectOptionSt);

    const type = Form.useWatch('type', { form, preserve: true });

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
            layout="vertical"
            form={form}
            name={formName}
            colon={false}
            labelWrap
            initialValues={{ ...initialValues }}
            onFinish={(data) => {
                const payload = { ...data, project_id: projectId };
                FormUtil.submit(endPoint, payload, method)
                    .then((data) => onChange(data, id))
                    .catch(FormUtil.setFormErrors(form, notification));
            }}
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
                name="type"
                label={labels.type}
                rules={[FormUtil.ruleRequired()]}
            >
                <SelectInput
                    block
                    disabled={!!id}
                    options={projectOption.task_field.type}
                />
            </Form.Item>
            {['SELECT', 'MULTIPLE_SELECT'].includes(type) ? (
                <Form.Item name="task_field_options" label={labels.task_field_options}>
                    <TaskFieldOptionTable />
                </Form.Item>
            ) : null}
        </Form>
    );
}

TaskFieldForm.displayName = formName;
TaskFieldForm.formName = formName;
