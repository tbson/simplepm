import * as React from 'react';
import { useRef, useEffect } from 'react';
import { Form, Input, Row, Col, InputNumber } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import ImgInput from 'component/common/form/ant/input/img_input';
import { urls, getLabels } from '../config';

const { TextArea } = Input;

const formName = 'WorkspaceForm';
const emptyRecord = {
    id: 0,
    tenant_id: null,
    title: '',
    description: '',
    order: 0,
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * WorkspaceForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function WorkspaceForm({ data, onChange }) {
    const inputRef = useRef(null);
    const [form] = Form.useForm();

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
            layout="vertical"
            initialValues={{ ...initialValues }}
            onFinish={(payload) =>
                FormUtil.submit(endPoint, payload, method)
                    .then((data) => onChange(data, id))
                    .catch(FormUtil.setFormErrors(form))
            }
        >
            <Row gutter={40}>
                <Col span={8}>
                    <Form.Item
                        name="avatar"
                        label={labels.avatar}
                    >
                        <ImgInput />
                    </Form.Item>
                </Col>
                <Col span={16}>
                    <Form.Item
                        name="title"
                        label={labels.title}
                        rules={[FormUtil.ruleRequired()]}
                    >
                        <Input ref={inputRef}/>
                    </Form.Item>

                    <Form.Item
                        name="description"
                        label={labels.description}
                    >
                        <TextArea />
                    </Form.Item>

                    <Form.Item
                        name="order"
                        label={labels.order}
                    >
                        <InputNumber className="full-width" />
                    </Form.Item>
                </Col>
            </Row>
        </Form>
    );
}

WorkspaceForm.displayName = formName;
WorkspaceForm.formName = formName;
