import * as React from "react";
import { t } from 'ttag';
import { Result, Button } from "antd";
import { HomeFilled } from "@ant-design/icons";
import { Link } from "react-router-dom";

export default function VerifyEmail() {
    return (
        <Result
            status="success"
            title={t`We have sent you an verification email`}
            subTitle={t`Please check your email and click on the link to verify your email`}
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

VerifyEmail.displayName = "VerifyEmail";
