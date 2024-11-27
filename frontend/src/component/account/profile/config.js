import { t } from "ttag";
import RequestUtil from "service/helper/request_util";

const urlMap = {
    base: {
        prefix: "account",
        endpoints: {
            crud: "",
            profile: "profile",
            password: "profile/password"
        }
    }
};
export const urls = RequestUtil.prefixMapValues(urlMap.base);

const headingTxt = t`Admin`;
const name = headingTxt.toLowerCase();
export const messages = {
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
};

export const emptyRecord = {
    id: 0,
    last_name: "",
    first_name: "",
    email: "",
    phone_number: "",
    groups: []
};

export const labels = {
    full_name: t`Fullname`,
    last_name: t`Lastname`,
    first_name: t`Firstname`,
    email: t`Email`,
    phone_number: t`Phone number`,
    is_active: t`Active`,
    groups: t`Groups`
};
