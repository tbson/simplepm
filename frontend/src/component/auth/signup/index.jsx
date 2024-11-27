import * as React from 'react';
import { useEffect } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { t } from 'ttag';
import { Row, Col, Card, Button, Divider } from 'antd';
import NavUtil from 'service/helper/nav_util';
import StorageUtil from 'service/helper/storage_util';
import LocaleSelect from 'component/common/locale_select.jsx';
import Form from './form';

const styles = {
    wrapper: {
        marginTop: 20
    }
};
export default function Signup() {
    const navigate = useNavigate();
    const navigateTo = NavUtil.navigateTo(navigate);

    useEffect(() => {
        StorageUtil.getUserInfo() && navigateTo();
    }, []);

    const handleSignup = () => {
        navigateTo("/verify-email");
    };

    return (
        <div>
            <div className="right content">
                <LocaleSelect />
            </div>
            <Row>
                <Col
                    xs={{ span: 24 }}
                    md={{ span: 12, offset: 6 }}
                    lg={{ span: 8, offset: 8 }}
                >
                    <Card title={t`Signup`} style={styles.wrapper}>
                        <Form onChange={handleSignup} />
                        <Divider plain>Already had account?</Divider>
                        <div className="center">
                            <Link to="/login">
                                <Button type="link">Login</Button>
                            </Link>
                        </div>
                    </Card>
                </Col>
            </Row>
        </div>
    );
}
