import { atom } from "jotai";
import LocaleUtil from "service/helper/locale_util";

export const localeSt = atom(LocaleUtil.getLocale());
