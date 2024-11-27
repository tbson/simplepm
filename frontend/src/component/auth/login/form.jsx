import * as React from 'react';
import { Button, Row, Col, Form, Input } from 'antd';
import { t } from 'ttag';
import { CheckOutlined } from '@ant-design/icons';
import StorageUtil from 'service/helper/storage_util';
import FormUtil from 'service/helper/form_util';
import { urls } from 'component/auth/config';

const formName = 'LoginForm';

export default function LoginForm({ onChange, children }) {
    const [form] = Form.useForm();
    const initialValues = {
        tenantUid: StorageUtil.getTenantUid()
    };

    const checkAuthUrl = (tenantUid) => {
        FormUtil.submit(`${urls.loginCheck}${tenantUid}`, {}, 'get')
            .then(() => {
                StorageUtil.setStorage('tenantUid', tenantUid);
                onChange(tenantUid);
            })
            .catch(FormUtil.setFormErrors(form));
    };

    return (
        <Form
            form={form}
            labelCol={{ span: 8 }}
            wrapperCol={{ span: 16 }}
            initialValues={{ ...initialValues }}
            onFinish={(payload) => {
                checkAuthUrl(payload.tenantUid);
            }}
        >
            <Form.Item
                name="tenantUid"
                label={t`Company code`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input autoFocus />
            </Form.Item>

            <br />
            <Row>
                <Col span={12}>{children}</Col>
                <Col span={12} className="right">
                    <Button type="primary" htmlType="submit" icon={<CheckOutlined />}>
                        {t`Login`}
                    </Button>
                </Col>
            </Row>
        </Form>
    );
}
LoginForm.displayName = formName;
LoginForm.formName = formName;
