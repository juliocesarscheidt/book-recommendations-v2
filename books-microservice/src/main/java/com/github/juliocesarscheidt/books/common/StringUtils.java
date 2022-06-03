package com.github.juliocesarscheidt.books.common;

public class StringUtils {

  public static String generateUuid(int size) {
    String result = "";
    String characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    for (int i = 0; i < size; i++) {
      result += characters.charAt((int) Math.floor(Math.random() * characters.length()));
    }
    return result;
  }
}
