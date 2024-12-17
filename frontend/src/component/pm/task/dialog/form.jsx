import * as React from 'react';
import { useRef, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { useAtomValue } from 'jotai';
import { Form, Input } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import SelectInput from 'component/common/form/ant/input/select_input';
import { taskOptionSt } from 'component/pm/task/state';
import { urls, getLabels } from '../config';

const { TextArea } = Input;

const formName = 'TaskForm';
const emptyRecord = {
    id: 0,
    title: '',
    description: '',
    feature_id: null
};

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
    const { project_id } = useParams();
    const inputRef = useRef(null);
    const [form] = Form.useForm();
    const taskOption = useAtomValue(taskOptionSt);

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
                FormUtil.submit(
                    endPoint,
                    { ...payload, project_id: parseInt(project_id) },
                    method
                )
                    .then((data) => onChange(data, id))
                    .catch(FormUtil.setFormErrors(form))
            }
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
                name="feature_id"
                label={labels.feature}
                rules={[FormUtil.ruleRequired()]}
            >
                <SelectInput block options={taskOption.feature} />
            </Form.Item>
        </Form>
    );
}

TaskForm.displayName = formName;
TaskForm.formName = formName;
