export default class NumberUtil {
    static isDigit(input) {
        return (
            (typeof input === 'string' && /^[0-9]$/.test(input)) ||
            (typeof input === 'number' && Number.isInteger(input))
        );
    }
}
