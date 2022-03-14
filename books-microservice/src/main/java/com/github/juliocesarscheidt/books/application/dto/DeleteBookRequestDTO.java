package com.github.juliocesarscheidt.books.application.dto;

public class DeleteBookRequestDTO extends BaseDTO {

  public String uuid;

  public String getUuid() {
    return uuid;
  }

  public void setUuid(String uuid) {
    this.uuid = uuid;
  }

  @Override
  public String toString() {
    return "DeleteBookRequestDTO [uuid=" + uuid + "]";
  }
}
