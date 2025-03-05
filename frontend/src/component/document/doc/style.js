export function getStyles(createStyles) {
    return createStyles(({ token, css }) => {
        console.log('background', token.colorBgLayout);
        console.log('border-bottom', token.colorBorder);
        return {
            layout: css`
                width: 100%;
                min-width: 1000px;
                height: 100%;
                border-radius: ${token.borderRadius}px;
                display: flex;
                background: ${token.colorBgContainer};
                font-family: AlibabaPuHuiTi, ${token.fontFamily}, sans-serif;

                .ant-prompts {
                    color: ${token.colorText};
                }
            `,
            menu: css`
                background: ${token.colorBgLayout}80;
                width: 25%;
                height: 100%;
                display: flex;
                flex-direction: column;
            `,
            conversations: css`
                padding: 0 5px;
                flex: 1;
                overflow-y: auto;
            `,
            chat: css`
                height: 100%;
                width: 100%;
                flex: 1;
                // max-width: 700px;
                padding: 0 !important;
                margin: 0 auto;
                box-sizing: border-box;
                display: flex;
                flex-direction: column;
                padding: ${token.paddingLG}px;
                gap: 16px;
            `,
            messages: css`
                flex: 1;
            `,
            placeholder: css`
                padding-top: 32px;
            `,
            sender: css`
                box-shadow: ${token.boxShadow};
            `,
            logo: css`
                display: flex;
                height: 72px;
                align-items: center;
                justify-content: start;
                padding: 0 24px;
                box-sizing: border-box;

                img {
                    width: 24px;
                    height: 24px;
                    display: inline-block;
                }

                span {
                    display: inline-block;
                    margin: 0 8px;
                    font-weight: bold;
                    color: ${token.colorText};
                    font-size: 16px;
                }
            `,
            addBtn: css`
                background: #1677ff0f;
                border: 1px solid #1677ff34;
                width: calc(100% - 24px);
                margin: 0 12px 24px 12px;
            `,
            document: css`
                width: 25%;
                height: 100%;
                display: flex;
                flex-direction: column;
                gap: 10px;
            `,
            documentRow: css`
                padding: 10px;
                padding-top: 0;
                padding-bottom: 0;
            `,
            chatHeading: css`
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 0 12px;
                height: 48px;
                font-weight: bold;
                background: ${token.colorBgLayout}80;
                border-bottom: 1px solid ${token.colorBorder};
            `
        };
    });
}
