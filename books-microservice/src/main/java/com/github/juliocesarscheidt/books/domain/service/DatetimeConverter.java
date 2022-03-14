package com.github.juliocesarscheidt.books.domain.service;

import java.sql.Timestamp;
import java.time.ZoneId;
import java.time.ZonedDateTime;
import java.util.Date;

public class DatetimeConverter {

  public static Date convertDateToGMT(Date date) {
    ZonedDateTime zdt = ZonedDateTime.ofInstant(date.toInstant(), ZoneId.of("GMT"));
    return Timestamp.valueOf(zdt.toLocalDateTime());
  }
}
