import * as React from "react";
import { t } from 'ttag';
import { Result, Button } from "antd";
import { HomeFilled } from "@ant-design/icons";
import { Link } from "react-router-dom";

export default function AuthError() {
    return (
        <Result
            status="error"
            title="Can not authenticate"
            subTitle={t`Sorry, there was a problem with authentication. Please contact the administrator.`}
            extra={
                <Link to="/">
                    {" "}
                    <Button type="primary" icon={<HomeFilled />}>
                        {t`Back to Home`}
                    </Button>
                </Link>
            }
        />
    );
}

AuthError.displayName = "AuthError";
