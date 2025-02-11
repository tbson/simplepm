import * as React from 'react';
import { useRef, useEffect } from 'react';
import { useParams, Link } from 'react-router';
import { App, Form, Input, Button, Row, Col } from 'antd';
import { createStyles } from 'antd-style';
import { LeftOutlined, CheckOutlined } from '@ant-design/icons';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import RichTextInput from 'component/common/form/ant/input/richtext_input';
import { getStyles } from './style';
import { urls, getLabels } from './config';

const formName = 'DocForm';
const emptyRecord = {
    id: 0,
    task_id: 0,
    title: '',
    content: {}
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * DocForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function DocForm({ data, onChange }) {
    let { taskId } = useParams();
    taskId = parseInt(taskId, 10);
    const { notification } = App.useApp();
    const inputRef = useRef(null);
    const [form] = Form.useForm();

    const useStyle = getStyles(createStyles);
    const { styles } = useStyle();
    const labels = getLabels();

    const initialValues = Util.isEmpty(data) ? emptyRecord : data;
    const { id } = initialValues;

    const endPoint = id ? `${urls.crud}${id}` : urls.crud;
    const method = id ? 'put' : 'post';

    useEffect(() => {
        setTimeout(() => {
            inputRef.current.focus({ cursor: 'end' });
        }, 100);
    }, []);

    return (
        <div>
            <div className={styles.chatHeading}>
                <div className="flex-item-remaining">
                    <Link to={`/pm/task/${taskId}`}>
                        <Button icon={<LeftOutlined />}>Back</Button>
                    </Link>
                </div>
                <div>
                    <Button
                        type="primary"
                        form={formName}
                        htmlType="submit"
                        icon={<CheckOutlined />}
                    >
                        Submit
                    </Button>
                </div>
            </div>
            <div className="content">
                <Form
                    form={form}
                    name={formName}
                    colon={false}
                    layout="vertical"
                    labelWrap
                    initialValues={{ ...initialValues }}
                    onFinish={(payload) => {
                        payload.task_id = taskId;
                        payload.type = 'DOC';
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

                    <Form.Item name="content" label={labels.content}>
                        <RichTextInput taskId={taskId} />
                    </Form.Item>
                </Form>
            </div>
        </div>
    );
}

DocForm.displayName = formName;
DocForm.formName = formName;
