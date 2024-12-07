import { format } from 'date-fns';

export const DATE_REABLE_FORMAT = 'dd/MM/yyyy';
export const DATE_ISO_FORMAT = 'yyyy-MM-dd';
export const DATETIME_ISO_FORMAT = "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'";

export default class DateUtil {
    /**
     * Format date to readable format.
     *
     * @param {Date} date
     * @returns {string}
     */
    static toReadableDate(date) {
        return format(date, DATE_REABLE_FORMAT);
    }

    /**
     * Format date to iso date format.
     *
     * @param {Date} date
     * @returns {string}
     */
    static toIsoDate(date) {
        try {
            return format(date, DATE_ISO_FORMAT);
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
        try {
            return format(date, DATETIME_ISO_FORMAT);
        } catch (e) {
            return null;
        }
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
