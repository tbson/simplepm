import * as React from 'react';
import { useEffect } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { t } from 'ttag';
import { Row, Col, Card, Button, Divider } from 'antd';
import Util from 'service/helper/util';
import NavUtil from 'service/helper/nav_util';
import StorageUtil from 'service/helper/storage_util';
import LocaleSelect from 'component/common/locale_select.jsx';
import Form from './form';

const styles = {
    wrapper: {
        marginTop: 20
    }
};
export default function Login() {
    const navigate = useNavigate();
    const navigateTo = NavUtil.navigateTo(navigate);

    useEffect(() => {
        StorageUtil.getUserInfo() && navigateTo();
    }, []);

    const handleLogin = (tenantUid) => {
        setTimeout(() => {
            Util.toggleGlobalLoading();
        }, 100);
        const ssoUrl = `/api/v1/account/auth/sso/login/${tenantUid}`;
        window.location.href = ssoUrl;
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
                    <Card title={t`Login`} style={styles.wrapper}>
                        <Form onChange={handleLogin} />
                        <Divider plain>Don't have account yet?</Divider>
                        <div className="center">
                            <Link to="/signup">
                                <Button type="link">Signup</Button>
                            </Link>
                        </div>
                    </Card>
                </Col>
            </Row>
        </div>
    );
}
