import React from 'react';

/**
    @param {Object} props
    @param {React.ReactNode} props.children
*/
export default function PageHeading({ children }) {
    return (
        <div
            style={{
                backgroundColor: '#f5f5f5',
                display: 'flex',
                alignItems: 'center',
                height: 40,
                fontWeight: 'bold',
                paddingLeft: 12
            }}
        >
            {children}
        </div>
    );
}

PageHeading.displayName = 'PageHeading';
