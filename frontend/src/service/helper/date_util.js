import { format } from "date-fns";

export default class DateUtil {
    /**
     * Format date to readable format.
     *
     * @param {Date} date
     * @returns {string}
     */
    static toReadableDate(date) {
        return format(date, "dd/MM/yyyy");
    }

    /**
     * Format date to iso date format.
     *
     * @param {Date} date
     * @returns {string}
     */
    static toIsoDate(date) {
        try {
            return format(date, "yyyy-MM-dd");
        } catch (e) {
            return null;
        }
    }

    /**
     * Format date to iso date time format.
     *
     * @param {Date} date
     * @returns {string}
     */
    static toIsoDateTime(date) {
        return format(date, "yyyy-MM-dd HH:mm:ss");
    }

    /**
     * Format string to date.
     *
     * @param {string} dateStr
     * @returns {Date}
     */
    static strToDate(dateStr) {
        return new Date(dateStr);
    }
}
