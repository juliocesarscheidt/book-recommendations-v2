package com.github.juliocesarscheidt.books.application.dto;

public class CreateBookResponseDTO extends BaseDTO {

  public String uuid;

  public CreateBookResponseDTO() {
  }

  public CreateBookResponseDTO(String uuid) {
    this.uuid = uuid;
  }

  public String getUuid() {
    return uuid;
  }

  public void setUuid(String uuid) {
    this.uuid = uuid;
  }

  @Override
  public String toString() {
    return "CreateBookResponseDTO [uuid=" + uuid + "]";
  }
}
