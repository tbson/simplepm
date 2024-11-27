import * as React from 'react';
import { useState, useEffect } from 'react';
import Img from 'component/common/display/img';

/**
 * ImgInput.
 *
 * @param {Object} props
 * @param {string} props.value
 * @param {function} props.onChange
 * @param {string} props.label
 * @returns {ReactElement}
 */
export default function ImgInput({ value, onChange }) {
    const [src, setSrc] = useState(value);

    useEffect(() => {
        setSrc(value);
    }, []);

    // Function to handle file input change
    const handleFileChange = (event) => {
        const file = event.target.files[0];
        if (file) {
            const blobUrl = URL.createObjectURL(file);
            console.log(blobUrl);
            setSrc(blobUrl); // Set the Blob URL as src for preview
            onChange(file); // Pass the file to onChange
            event.target.value = null; // Reset the file input for re-selection
        }
    };

    // Hidden file input to allow image selection
    const handleClick = () => {
        document.getElementById('fileInput').click();
    };

    return (
        <div
            className="pointer"
            onClick={handleClick} // Trigger the file input on click
        >
            <Img src={src} width={150} height={150} preview={false} />
            <input
                id="fileInput"
                type="file"
                accept="image/*"
                style={{ display: 'none' }}
                onChange={handleFileChange} // Handle the file selection
            />
        </div>
    );
}
