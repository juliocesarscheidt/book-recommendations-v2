package com.github.juliocesarscheidt.books.application.usecase;

import java.math.BigInteger;
import java.sql.Timestamp;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.time.ZonedDateTime;
import java.util.Date;
import java.util.Random;

public class BaseUseCase {

  public static Timestamp getCurrentTimestampUTC() {
    ZonedDateTime zdt = ZonedDateTime.of(LocalDateTime.now(), ZoneId.of("CET"));
    return Timestamp.valueOf(zdt.toLocalDateTime());
  }
  
  public static String generateUuid() {
	  String randTimeHex = new BigInteger(String.valueOf(new Date().getTime()).getBytes()).toString(16).substring(0, 12);

	  Random r = new Random();
	  String code = "";
	  for (int i = 0; i < 12; i ++) {
		  char c = (char)(r.nextInt(26) + 'a');
		  code = code + c + randTimeHex.charAt(i);
	  }
	  String uuid = code.substring(0, 24);
	  System.out.println("uuid :: " + uuid);
	  
	  return uuid;
  }
}
