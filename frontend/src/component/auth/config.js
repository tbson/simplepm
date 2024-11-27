import RequestUtil from "service/helper/request_util";

const urlMap = {
    base: {
        prefix: "account/auth/sso",
        endpoints: {
            loginCheck: "login/check",
        }
    },
    signup: {
        prefix: "account/signup-tenant",
        endpoints: {
            signup: "",
        }
    }
};

export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const signupUrls = RequestUtil.prefixMapValues(urlMap.signup);

const headingTxt = "Hồ sơ";
export const messages = {
    heading: headingTxt
};
