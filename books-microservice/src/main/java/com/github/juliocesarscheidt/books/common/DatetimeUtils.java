package com.github.juliocesarscheidt.books.common;

import java.sql.Timestamp;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.time.ZonedDateTime;
import java.util.Date;

public class DatetimeUtils {

  public static Date convertDateToGMT(Date date) {
    ZonedDateTime zdt = ZonedDateTime.ofInstant(date.toInstant(), ZoneId.of("GMT"));
    return Timestamp.valueOf(zdt.toLocalDateTime());
  }

  public static Timestamp getCurrentTimestampUTC() {
    ZonedDateTime zdt = ZonedDateTime.of(LocalDateTime.now(), ZoneId.of("CET"));
    return Timestamp.valueOf(zdt.toLocalDateTime());
  }
}
