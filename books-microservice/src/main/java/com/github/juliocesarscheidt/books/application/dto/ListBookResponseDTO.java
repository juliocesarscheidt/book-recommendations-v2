package com.github.juliocesarscheidt.books.application.dto;

import java.util.List;

public class ListBookResponseDTO extends BaseDTO {

  public List<GetBookResponseDTO> books;

  public ListBookResponseDTO() {
  }

  public ListBookResponseDTO(List<GetBookResponseDTO> books) {
    this.books = books;
  }

  public List<GetBookResponseDTO> getBooks() {
    return books;
  }

  public void setBooks(List<GetBookResponseDTO> books) {
    this.books = books;
  }

  @Override
  public String toString() {
    return "ListBookResponseDTO [books=" + books + "]";
  }
}
