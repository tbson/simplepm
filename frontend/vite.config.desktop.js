import path from 'path';
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';

const host = process.env.TAURI_DEV_HOST;

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
    // Vite options tailored for Tauri development and only applied in `tauri dev` or `tauri build`
    //
    // 1. prevent vite from obscuring rust errors
    clearScreen: false,
    // 2. tauri expects a fixed port, fail if that port is not available
    server: {
        port: 1420,
        strictPort: true,
        host: host || false,
        hmr: host
            ? {
                  protocol: 'ws',
                  host,
                  port: 1421
              }
            : undefined,
        watch: {
            // 3. tell vite to ignore watching `src-tauri`
            ignored: ['**/src-tauri/**']
        }
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
