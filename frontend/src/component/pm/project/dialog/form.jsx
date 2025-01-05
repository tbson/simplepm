import * as React from 'react';
import { useRef, useEffect } from 'react';
import { useAtomValue } from 'jotai';
import { App, Form, Input } from 'antd';
import Util from 'service/helper/util';
import RequestUtil from 'service/helper/request_util';
import FormUtil from 'service/helper/form_util';
import SelectInput from 'component/common/form/ant/input/select_input';
import ImgInput from 'component/common/form/ant/input/img_input';
import { projectOptionSt } from 'component/pm/project/state';
import { urls, getLabels } from '../config';

const { TextArea } = Input;

const dateFields = [];

const formName = 'ProjectForm';
const emptyRecord = {
    id: 0,
    workspace_id: null,
    title: '',
    description: '',
    avatar: '',
    layout: 'TABLE',
    status: 'ACTIVE',
    order: 0,
    finished_at: '',
    git_repo: ''
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
    const { notification } = App.useApp();
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
            onFinish={(data) => {
                const payload = RequestUtil.formatPayloadDate(data, dateFields);
                if (!id) {
                    payload.layout = 'KANBAN';
                    payload.status = 'ACTIVE';
                }
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

            <Form.Item name="git_repo" label={labels.git_repo}>
                <Input />
            </Form.Item>

            <Form.Item name="avatar" label={labels.avatar}>
                <ImgInput />
            </Form.Item>

            {/*
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
            */}
            {id ? (
                <Form.Item
                    name="status"
                    label={labels.status}
                    rules={[FormUtil.ruleRequired()]}
                >
                    <SelectInput block options={projectOption.status} />
                </Form.Item>
            ) : null}
        </Form>
    );
}

ProjectForm.displayName = formName;
ProjectForm.formName = formName;
