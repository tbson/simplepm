import Paragraph from '@editorjs/paragraph';
import Table from '@editorjs/table';
import List from '@editorjs/list';
import Code from '@editorjs/code';
import LinkTool from '@editorjs/link';
import Image from '@editorjs/image';
import Header from '@editorjs/header';
import Quote from '@editorjs/quote';
import Marker from '@editorjs/marker';
import Delimiter from '@editorjs/delimiter';
import InlineCode from '@editorjs/inline-code';
import RequestUtil from 'service/helper/request_util';

export function getTools(taskId) {
    return {
        paragraph: { class: Paragraph, inlineToolbar: true },
        header: Header,
        image: {
            class: Image,
            config: {
                uploader: {
                    uploadByFile(file) {
                        const url = `/document/docattachment/?task_id=${taskId}`;
                        const payload = { file, task_id: taskId };
                        return RequestUtil.apiCall(url, payload, 'post').then(
                            (resp) => {
                                return {
                                    success: 1,
                                    file: {
                                        url: resp.data.file_url
                                    }
                                };
                            }
                        );
                    }
                }
            }
        },
        linkTool: LinkTool,
        code: Code,
        list: { class: List, inlineToolbar: true },
        table: Table,
        quote: Quote,
        marker: Marker,
        delimiter: Delimiter,
        inlineCode: InlineCode
    };
}
