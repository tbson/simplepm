import * as React from 'react';
import { useRef, useEffect } from 'react';
import { useAtomValue } from 'jotai';
import { Form, Input } from 'antd';
import { useParams } from 'react-router-dom';
import Util from 'service/helper/util';
import FormUtil from 'service/helper/form_util';
import TransferInput from 'component/common/form/ant/input/transfer_input.jsx';
import { roleTransferSt } from 'component/account/role/state';
import { urls, getLabels } from '../config';

const formName = 'RoleForm';
const emptyRecord = {
    id: 0,
    tenant_id: null,
    title: '',
    pems: []
};

/**
 * @callback FormCallback
 *
 * @param {Object} data
 * @param {number} id
 */

/**
 * RoleForm.
 *
 * @param {Object} props
 * @param {Object} props.data
 * @param {FormCallback} props.onChange
 */
export default function RoleForm({ data, onChange }) {
    const { tenant_id } = useParams();
    const inputRef = useRef(null);
    const [form] = Form.useForm();
    const roleTransfer = useAtomValue(roleTransferSt);

    const labels = getLabels();

    const initialValues = Util.isEmpty(data) ? emptyRecord : data;
    const { id } = initialValues;

    let endPoint = id ? `${urls.crud}${id}` : urls.crud;
    if (tenant_id) {
        endPoint += `?tenant_id=${tenant_id}`;
    }
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
            onFinish={(payload) => {
                FormUtil.submit(endPoint, payload, method)
                    .then((data) => onChange(data, id))
                    .catch(FormUtil.setFormErrors(form));
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
                name="pem_ids"
                label={labels.pem_ids}
                rules={[FormUtil.ruleRequired()]}
            >
                <TransferInput options={roleTransfer.pem} />
            </Form.Item>
        </Form>
    );
}

RoleForm.displayName = formName;
RoleForm.formName = formName;
