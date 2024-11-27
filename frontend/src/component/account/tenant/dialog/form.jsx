import * as React from 'react';
import { useRef, useEffect } from 'react';
import { useAtomValue } from 'jotai';
import { Form, Input, Row, Col } from 'antd';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import SelectInput from 'component/common/form/ant/input/select_input';
import ImgInput from 'component/common/form/ant/input/img_input';
import { tenantOptionSt } from 'component/account/tenant/state';
import { urls, getLabels } from '../config';

const formName = 'TenantForm';
const emptyRecord = {
    id: 0,
    auth_client_id: null,
    uid: '',
    title: ''
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * TenantForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function TenantForm({ data, onChange }) {
    const inputRef = useRef(null);
    const [form] = Form.useForm();
    const tenantOption = useAtomValue(tenantOptionSt);

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
                        name="uid"
                        label={labels.uid}
                        rules={[FormUtil.ruleRequired()]}
                    >
                        <Input ref={inputRef} />
                    </Form.Item>

                    <Form.Item
                        name="title"
                        label={labels.title}
                        rules={[FormUtil.ruleRequired()]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        name="auth_client_id"
                        label={labels.auth_client_id}
                        rules={[FormUtil.ruleRequired()]}
                    >
                        <SelectInput block options={tenantOption.auth_client} />
                    </Form.Item>
                </Col>
            </Row>
        </Form>
    );
}

TenantForm.displayName = formName;
TenantForm.formName = formName;
