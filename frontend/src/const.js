export const ENV = import.meta.env;

export const APP_NAMESPACE = ENV.VITE_NAMESPACE;
export const DEV_MODE = parseInt(ENV.VITE_DEV_MODE);
export const CENTRIFUGO_SUBSCRIPTION_TOKEN_ENDPOINT =
    ENV.VITE_CENTRIFUGO_SUBSCRIPTION_TOKEN_ENDPOINT;
export const CENTRIFUGO_SOCKET_ENDPOINT = ENV.VITE_CENTRIFUGO_SOCKET_ENDPOINT;

//export const PROTOCOL = window.location.protocol + '//';
export const PROTOCOL = 'https://';
export const DOMAIN = window.location.host;
// export const DOMAIN = "simplepm.test";
export const API_PREFIX = '/api/v1/';
export const CLIENT_TYPE = 'web';

export const LOCAL_STORAGE_PREFIX = APP_NAMESPACE;

export const LOGO_TEXT = ENV.VITE_LOGO_TEXT || 'LOGO';
