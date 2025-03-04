import path from 'path';
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';
/*
const ReactCompilerConfig = {
    optimize: true,
    runtime: 'automatic' // e.g., 'classic' or 'automatic' for JSX runtime
};
*/
// https://vitejs.dev/config/
export default defineConfig({
    build: {
        target: 'esnext'
    },
    resolve: {
        alias: {
            src: path.resolve(__dirname, './src'),
            service: path.resolve(__dirname, './src/service'),
            component: path.resolve(__dirname, './src/component'),
            style: path.resolve(__dirname, './src/style')
        }
    },
    server: {
        host: '0.0.0.0',
        port: 3000,
        hmr: {
            host: 'simplepm.test',
            clientPort: 443,
            protocol: 'wss'
        },
        allowedHosts: ['simplepm.test', 'frontend']
    },
    plugins: [
        react({
            fastRefresh: process.env.NODE_ENV !== 'test'
            /*
            babel: {
                plugins: [['babel-plugin-react-compiler', ReactCompilerConfig]]
            }
            */
        })
    ],
    test: {
        globals: true,
        environment: 'jsdom',
        setupFiles: ['./src/vitest.setup.js']
    }
});
