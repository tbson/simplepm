import * as React from "react";
import { t } from 'ttag';
import { Result, Button } from "antd";
import { HomeFilled } from "@ant-design/icons";
import { Link } from "react-router-dom";

export default function NotMatch() {
    return (
        <Result
            status="404"
            title="404"
            subTitle={t`Sorry, the page you visited does not exist.`}
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

NotMatch.displayName = "NotMatch";
