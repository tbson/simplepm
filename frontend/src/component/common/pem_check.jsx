import PemUtil from "service/helper/pem_util";

export default function PemCheck({ pem_group, pem, children }) {
    if (PemUtil.hasPermit(pem_group, pem)) {
        return children;
    }
    return null;
}
