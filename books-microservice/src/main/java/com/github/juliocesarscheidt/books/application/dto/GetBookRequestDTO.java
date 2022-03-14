package com.github.juliocesarscheidt.books.application.dto;

public class GetBookRequestDTO extends BaseDTO {

  public String uuid;

  public String getUuid() {
    return uuid;
  }

  public void setUuid(String uuid) {
    this.uuid = uuid;
  }

  @Override
  public String toString() {
    return "GetBookRequestDTO [uuid=" + uuid + "]";
  }
}
