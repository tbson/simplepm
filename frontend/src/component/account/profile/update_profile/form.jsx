import * as React from 'react';
import { useEffect, useRef } from 'react';
import { t } from 'ttag';
import { Form, Input, Row, Col } from 'antd';
import FormUtil from 'service/helper/form_util';
import ImgInput from 'component/common/form/ant/input/img_input';
import { urls } from '../config';

const formName = 'UpdateProfileForm';

export default function UpdateProfileForm({ data, onChange }) {
    const inputRef = useRef(null);
    const [form] = Form.useForm();

    useEffect(() => {
        inputRef.current.focus({ cursor: 'end' });
    }, []);

    const formAttrs = {
        avatar: {
            name: 'avatar',
            label: t`Avatar`
        },
        mobile: {
            name: 'mobile',
            label: t`Mobile`
        },
        first_name: {
            name: 'first_name',
            label: t`First Name`,
            rules: [FormUtil.ruleRequired()]
        },
        last_name: {
            name: 'last_name',
            label: t`Last Name`,
            rules: [FormUtil.ruleRequired()]
        }
    };

    return (
        <Form
            form={form}
            name={formName}
            layout="vertical"
            initialValues={{ ...data }}
            onFinish={(payload) => {
                console.log('payload', payload);
                FormUtil.submit(urls.profile, payload, 'put')
                    .then((data) => onChange(data))
                    .catch(FormUtil.setFormErrors(form));
            }}
        >
            <Row>
                <Col span={8}>
                    <Form.Item {...formAttrs.avatar}>
                        <ImgInput />
                    </Form.Item>
                </Col>
                <Col span={16}>
                    <Form.Item {...formAttrs.mobile}>
                        <Input ref={inputRef} />
                    </Form.Item>
                    <Form.Item {...formAttrs.first_name}>
                        <Input />
                    </Form.Item>
                    <Form.Item {...formAttrs.last_name}>
                        <Input />
                    </Form.Item>
                </Col>
            </Row>
        </Form>
    );
}

UpdateProfileForm.displayName = formName;
UpdateProfileForm.formName = formName;
