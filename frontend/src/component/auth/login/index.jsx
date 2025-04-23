import React, { useEffect, useCallback } from 'react';
import { useNavigate, Link } from 'react-router';
import { t } from 'ttag';
import { Row, Col, Card, Button, Divider } from 'antd';
import StorageUtil from 'service/helper/storage_util';
import NavUtil from 'service/helper/nav_util';
import LocaleSelect from 'component/common/locale_select.jsx';
import ResetPwdDialog from 'component/auth/reset_pwd';
import LoginForm from './form';

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

    const handleLogin = ({auth, next}) => {
        StorageUtil.setStorage('auth', auth);
        navigateTo(next || '/');
    };

    const handleOpenResetPwd = useCallback(() => {
        ResetPwdDialog.toggle(true);
    }, []);

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
                        <LoginForm onChange={handleLogin}>
                            <Button
                                color="primary"
                                variant="link"
                                onClick={handleOpenResetPwd}
                            >
                                {t`Forgot password`}
                            </Button>
                        </LoginForm>
                        <Divider plain>Don’t have an account yet?</Divider>
                        <div className="center">
                            <Link to="/signup">
                                <Button type="link">
                                    Register a new team or company
                                </Button>
                            </Link>
                        </div>
                    </Card>
                </Col>
            </Row>
            <ResetPwdDialog />
        </div>
    );
}
