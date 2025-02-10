// tools.js
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

export const EDITOR_JS_TOOLS = {
    paragraph: { class: Paragraph, inlineToolbar: true },
    header: Header,
    image: {
        class: Image,
        config: {
            endpoints: {
                byFile: 'http://localhost:8008/uploadFile',
                byUrl: 'http://localhost:8008/fetchUrl'
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
