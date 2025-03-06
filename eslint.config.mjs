// eslint.config.mjs
import reactPlugin from 'eslint-plugin-react';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import js from '@eslint/js';

export default [
    {
        ignores: [
            'backend',
            'frontend/dist',
            'frontend/vite.config.js',
            'frontend/vite.config.web.js',
            'frontend/vite.config.desktop.js'
        ]
    },
    {
        // Replace "env" with "languageOptions.globals"
        languageOptions: {
            ecmaVersion: 2022,
            sourceType: 'module',
            parserOptions: {
                ecmaFeatures: {
                    jsx: true
                }
            },
            // Define globals for browser environment
            globals: {
                window: 'readonly',
                document: 'readonly',
                navigator: 'readonly',
                console: 'readonly',
                setTimeout: 'readonly',
                clearTimeout: 'readonly',
                setInterval: 'readonly',
                clearInterval: 'readonly',
                fetch: 'readonly',
                FormData: 'readonly',
                localStorage: 'readonly',
                Blob: 'readonly',
                URLSearchParams: 'readonly',
                CustomEvent: 'readonly',
                // Add any other globals you need
            }
        },
        plugins: {
            react: reactPlugin,
            'react-hooks': reactHooks,
            'react-refresh': reactRefresh
        },
        rules: {
            ...js.configs.recommended.rules,
            ...reactHooks.configs.recommended.rules,
            'react-refresh/only-export-components': [
                'warn',
                { allowConstantExport: true }
            ],
            'react/prop-types': 'off',
            'react/react-in-jsx-scope': 'off',
            'no-unused-vars': [
                'warn',
                {
                    argsIgnorePattern: '^_',
                    varsIgnorePattern: '^_',
                    caughtErrorsIgnorePattern: '^_'
                }
            ]
        }
    }
]
