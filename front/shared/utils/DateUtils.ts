import { DateTime } from "luxon";

export class DateUtils {
    public static toDateTime(date: Date): DateTime {
        return DateTime.fromJSDate(date);
    }

    public static formatDate = (date: Date) => {
        return DateTime.fromJSDate(new Date(date)).toFormat('dd.MM.yy');
    }

    public static dateTimeToString = (date: Date) => {
        return DateTime.fromJSDate(new Date(date)).toFormat('dd.MM.yy HH:mm');
    };

    public static currentDate = (): string => {
        return DateTime.fromJSDate(new Date(), ).toFormat('dd.MM.yyyy')
    }
}

