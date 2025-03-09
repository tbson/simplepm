import React from 'react';
import { t } from 'ttag';
import { App, Button, Row, Col, Form, Input } from 'antd';
import { ArrowRightOutlined } from '@ant-design/icons';
import StorageUtil from 'service/helper/storage_util';
import FormUtil from 'service/helper/form_util';
import { urls } from 'component/auth/config';

const formName = 'LoginForm';

export default function LoginForm({ onChange, children }) {
    const { notification } = App.useApp();
    const [form] = Form.useForm();
    const initialValues = {
        username: '',
        password: ''
    };

    const checkAuthUrl = (tenantUid) => {
        FormUtil.submit(`${urls.loginCheck}${tenantUid}`, {}, 'get')
            .then(() => {
                StorageUtil.setStorage('tenantUid', tenantUid);
                onChange(tenantUid);
            })
            .catch(FormUtil.setFormErrors(form, notification));
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
                name="username"
                label={t`Username`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input autoFocus />
            </Form.Item>

            <Form.Item
                name="password"
                label={t`Password`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input type="password" />
            </Form.Item>

            <br />
            <Row>
                <Col span={12}>{children}</Col>
                <Col span={12} className="right">
                    <Button
                        type="primary"
                        htmlType="submit"
                        icon={<ArrowRightOutlined />}
                    >
                        {t`Process`}
                    </Button>
                </Col>
            </Row>
        </Form>
    );
}
LoginForm.displayName = formName;
LoginForm.formName = formName;
