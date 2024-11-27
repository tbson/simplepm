import * as React from 'react';
import { useEffect, useRef } from 'react';
import { t } from 'ttag';
import { Row, Col, Form, Input } from 'antd';
import FormUtil from 'service/helper/form_util';
import { urls } from '../config';

const formName = 'ChangePwdForm';

const initValues = {
    password: '',
    password_confirm: ''
};

export default function ChangePwdForm({ onChange }) {
    const inputRef = useRef(null);
    const [form] = Form.useForm();

    useEffect(() => {
        inputRef.current.focus({ cursor: 'end' });
    }, []);

    const formAttrs = {
        password: {
            name: 'password',
            label: t`New password`,
            rules: [FormUtil.ruleRequired()]
        },
        password_confirm: {
            name: 'password_confirm',
            label: t`Confirm new password`,
            rules: [FormUtil.ruleRequired()]
        }
    };

    return (
        <Form
            layout="vertical"
            form={form}
            name={formName}
            initialValues={{ ...initValues }}
            onFinish={(payload) =>
                FormUtil.submit(urls.password, payload, 'put')
                    .then((data) => onChange(data))
                    .catch(FormUtil.setFormErrors(form))
            }
        >
            <Row gutter={24}>
                <Col span={12}>
                    <Form.Item {...formAttrs.password}>
                        <Input ref={inputRef} type="password" />
                    </Form.Item>
                </Col>
                <Col span={12}>
                    <Form.Item {...formAttrs.password_confirm}>
                        <Input type="password" />
                    </Form.Item>
                </Col>
            </Row>
        </Form>
    );
}

ChangePwdForm.displayName = formName;
ChangePwdForm.formName = formName;
