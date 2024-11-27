import * as React from 'react';
import { Button, Row, Col, Form, Input } from 'antd';
import { t } from 'ttag';
import { CheckOutlined } from '@ant-design/icons';
import FormUtil from 'service/helper/form_util';
import { signupUrls } from 'component/auth/config';

const formName = 'SignupForm';

export default function SignupForm({ onChange, children }) {
    const [form] = Form.useForm();
    const initialValues = {
        uid: '',
        title: '',
        email: '',
        mobile: null,
        first_name: '',
        last_name: '',
        password: ''
    };

    const handleSignup = (payload) => {
        FormUtil.submit(`${signupUrls.signup}`, payload, 'post')
            .then(() => {
                onChange(payload);
            })
            .catch(FormUtil.setFormErrors(form))
    };

    return (
        <Form
            form={form}
            labelCol={{ span: 8 }}
            wrapperCol={{ span: 16 }}
            initialValues={{ ...initialValues }}
            onFinish={(payload) => {
                handleSignup(payload);
            }}
        >
            <Form.Item
                name="uid"
                label={t`Company code`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input autoFocus />
            </Form.Item>

            <Form.Item
                name="title"
                label={t`Company name`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                name="email"
                label={t`Email`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                name="mobile"
                label={t`Mobile`}
            >
                <Input />
            </Form.Item>

            <Form.Item
                name="first_name"
                label={t`First name`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                name="last_name"
                label={t`Last name`}
                rules={[FormUtil.ruleRequired()]}
            >
                <Input />
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
                    <Button type="primary" htmlType="submit" icon={<CheckOutlined />}>
                        {t`Signup`}
                    </Button>
                </Col>
            </Row>
        </Form>
    );
}
SignupForm.displayName = formName;
SignupForm.formName = formName;
