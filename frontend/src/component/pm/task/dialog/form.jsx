import * as React from 'react';
import { useRef, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { useAtomValue } from 'jotai';
import { App, Form, Input, InputNumber } from 'antd';
import Util from 'service/helper/util';
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
export default function TaskForm({ data, onChange }) {
    const { notification } = App.useApp();
    const { project_id } = useParams();
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
    if (!initialValues.id && taskOption.feature.length > 0) {
        initialValues.feature_id = taskOption.feature[0].value;
    } else {
        initialValues.feature_id = data.feature.id;
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
        if (!value) return '';
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
                const fieldIdStr = key.replace('EXT_', '');
                const fieldId = parseInt(fieldIdStr);
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
                FormUtil.submit(
                    endPoint,
                    { ...data, project_id: parseInt(project_id) },
                    method
                )
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

            <Form.Item
                name="feature_id"
                label={labels.feature}
                rules={[FormUtil.ruleRequired()]}
            >
                <SelectInput block options={taskOption.feature} />
            </Form.Item>

            <Form.Item name="description" label={labels.description}>
                <TextArea />
            </Form.Item>

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
