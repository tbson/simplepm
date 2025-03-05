import React from 'react';
import { useParams, Link } from 'react-router';
import { Button } from 'antd';
import { LeftOutlined, EditOutlined } from '@ant-design/icons';
import { createStyles } from 'antd-style';
import RichTextInput from 'component/common/form/ant/input/richtext_input';
import { getStyles } from './style';

export default function DocView({ data, toggleMode }) {
    let { taskId } = useParams();
    taskId = parseInt(taskId, 10);
    const useStyle = getStyles(createStyles);
    const { styles } = useStyle();
    return (
        <div>
            <div className={styles.chatHeading}>
                <div className="flex-item-remaining">
                    <Link to={`/pm/task/${taskId}`}>
                        <Button icon={<LeftOutlined />}>Back</Button>
                    </Link>
                </div>
                <div>
                    <Button
                        icon={<EditOutlined />}
                        onClick={() => {
                            toggleMode();
                        }}
                    >
                        Edit
                    </Button>
                </div>
            </div>
            <div className="content">
                <h1>{data.title}</h1>
                <RichTextInput value={data.content} disabled />
            </div>
        </div>
    );
}
