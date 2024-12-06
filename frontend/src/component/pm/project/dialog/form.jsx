import * as React from 'react';
import { useRef, useEffect } from 'react';
import { useAtomValue } from 'jotai';
import { Form, Input, InputNumber } from 'antd';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import FormUtil from 'service/helper/form_util';
import SelectInput from 'component/common/form/ant/input/select_input';
import DateInput from 'component/common/form/ant/input/date_input';
import { projectOptionSt } from 'component/pm/project/state';
import { urls, getLabels } from '../config';

const { TextArea } = Input;

const dateFields = ['start_date', 'target_date'];

const formName = 'ProjectForm';
const emptyRecord = {
    id: 0,
    workspace_id: null,
    title: '',
    description: '',
    avatar: '',
    layout: 'TABLE',
    order: 0,
    start_date: '',
    target_date: '',
    finished_at: ''
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * ProjectForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function ProjectForm({ data, onChange }) {
    const inputRef = useRef(null);
    const [form] = Form.useForm();
    const projectOption = useAtomValue(projectOptionSt);

    const labels = getLabels();

    const initialValues = Util.isEmpty(data)
        ? emptyRecord
        : RequestUtil.formatResponseDate(data, dateFields);
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
            onFinish={(payload) => {
                FormUtil.submit(endPoint, payload, method)
                    .then((data) => onChange(data, id))
                    .catch(FormUtil.setFormErrors(form));
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

            <Form.Item name="workspace_id" label={labels.workspace_id}>
                <SelectInput
                    block
                    options={FormUtil.addOptional(projectOption.workspace)}
                />
            </Form.Item>

            <Form.Item
                name="layout"
                label={labels.layout}
                rules={[FormUtil.ruleRequired()]}
            >
                <SelectInput block options={projectOption.layout} />
            </Form.Item>

            <Form.Item name="start_date" label={labels.start_date}>
                <DateInput />
            </Form.Item>

            <Form.Item name="target_date" label={labels.target_date}>
                <DateInput />
            </Form.Item>

            <Form.Item name="order" label={labels.order}>
                <InputNumber />
            </Form.Item>
        </Form>
    );
}

ProjectForm.displayName = formName;
ProjectForm.formName = formName;
