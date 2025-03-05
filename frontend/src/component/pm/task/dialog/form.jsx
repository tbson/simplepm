import React, { useRef, useEffect } from 'react';
import { useAtomValue } from 'jotai';
import { App, Space, Button, Form, Input, InputNumber } from 'antd';
import { PlusOutlined, MinusCircleOutlined } from '@ant-design/icons';
import DateUtil from 'service/helper/date_util';
import FormUtil from 'service/helper/form_util';
import SelectInput from 'component/common/form/ant/input/select_input';
import DateInput from 'component/common/form/ant/input/date_input';
import { taskOptionSt } from 'component/pm/task/state';
import { urls, getLabels } from '../config';

const { TextArea } = Input;

const formName = 'TaskForm';

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * TaskForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function TaskForm({ projectId, data, onChange }) {
    const { notification } = App.useApp();
    const inputRef = useRef(null);
    const [form] = Form.useForm();
    const taskOption = useAtomValue(taskOptionSt);
    const taskField = taskOption.task_field;
    const fieldTypeMap = taskField.reduce((acc, field) => {
        acc[field.value] = field.type;
        return acc;
    }, {});
    const statusField = taskField.find((field) => field.is_status);

    const labels = getLabels();

    const initialValues = { ...data };
    if (initialValues) {
        initialValues[`EXT_${statusField.value}`] = initialValues.status;
        delete initialValues.status;
    }

    for (const field of data?.task_fields || []) {
        const key = `EXT_${field.task_field_id}`;
        initialValues[key] = FormUtil.parseFieldValue(field.value, field.type);
    }
    const { id } = initialValues;

    const endPoint = id ? `${urls.crud}${id}` : urls.crud;
    const method = id ? 'put' : 'post';

    useEffect(() => {
        inputRef.current.focus({ cursor: 'end' });
    }, []);

    const renderDynamicField = (field) => {
        if (field.type === 'SELECT') {
            return <SelectInput block options={field.options} />;
        }
        if (field.type === 'MULTIPLE_SELECT') {
            return <SelectInput block mode="multiple" options={field.options} />;
        }
        if (field.type === 'DATE') {
            return <DateInput />;
        }
        if (field.type === 'NUMBER') {
            return <InputNumber className="full-width" />;
        }
        return <TextArea />;
    };

    const processFieldValue = (value, type) => {
        if (!value) {
            return '';
        }
        if (type === 'DATE') {
            return DateUtil.toIsoDate(value);
        }
        if (type === 'MULTIPLE_SELECT') {
            return value.join(',');
        }
        return `${value}`;
    };

    const processPayload = (payload) => {
        const data = { task_fields: [] };
        Object.entries(payload).forEach(([key, value]) => {
            if (key.startsWith('EXT_')) {
                const fieldId = parseInt(key.replace('EXT_', ''));
                const type = fieldTypeMap[fieldId];
                const field = {
                    task_field_id: fieldId,
                    value: processFieldValue(value, type)
                };
                data.task_fields.push(field);
            } else {
                data[key] = value;
            }
        });
        return data;
    };

    return (
        <Form
            form={form}
            name={formName}
            layout="vertical"
            initialValues={{ ...initialValues }}
            onFinish={(payload) => {
                const data = processPayload(payload);
                FormUtil.submit(endPoint, { ...data, project_id: projectId }, method)
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

            <Form.List name="task_users">
                {(fields, { add, remove }) => (
                    <div>
                        <div style={{ marginBottom: 8 }}>
                            <label>Users</label>
                        </div>
                        {fields.map(({ key, name, ...restField }) => (
                            <Space
                                key={key}
                                style={{
                                    display: 'flex',
                                    marginBottom: 8
                                }}
                                align="baseline"
                            >
                                <Form.Item
                                    {...restField}
                                    style={{ width: '200px' }}
                                    name={[name, 'user_id']}
                                    rules={[
                                        {
                                            required: true,
                                            message: 'Missing user'
                                        }
                                    ]}
                                >
                                    <SelectInput
                                        block
                                        placeholder="User"
                                        options={taskOption.user}
                                    />
                                </Form.Item>

                                <Form.Item
                                    {...restField}
                                    style={{ width: '200px' }}
                                    name={[name, 'git_branch']}
                                >
                                    <Input placeholder="Git branch" />
                                </Form.Item>
                                <MinusCircleOutlined onClick={() => remove(name)} />
                            </Space>
                        ))}
                        <Button
                            type="dashed"
                            onClick={() => add()}
                            block
                            icon={<PlusOutlined />}
                        >
                            Add user
                        </Button>
                    </div>
                )}
            </Form.List>
            <br />
            {taskField.map((field) => {
                return (
                    <Form.Item
                        key={field.value}
                        name={`EXT_${field.value}`}
                        label={field.label}
                        rules={field.is_status ? [FormUtil.ruleRequired()] : []}
                    >
                        {renderDynamicField(field)}
                    </Form.Item>
                );
            })}
        </Form>
    );
}

TaskForm.displayName = formName;
TaskForm.formName = formName;
