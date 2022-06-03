package com.github.juliocesarscheidt.books.application.dto;

import javax.validation.constraints.NotBlank;

public class UpdateBookRequestDTO extends BaseDTO {

  @NotBlank
  public String uuid;

  public String title;

  public String author;

  public String genre;

  public String image;

  public String getUuid() {
    return uuid;
  }

  public void setUuid(String uuid) {
    this.uuid = uuid;
  }

  public String getTitle() {
    return title;
  }

  public void setTitle(String title) {
    this.title = title;
  }

  public String getAuthor() {
    return author;
  }

  public void setAuthor(String author) {
    this.author = author;
  }

  public String getGenre() {
    return genre;
  }

  public void setGenre(String genre) {
    this.genre = genre;
  }

  public String getImage() {
    return image;
  }

  public void setImage(String image) {
    this.image = image;
  }

  @Override
  public String toString() {
    return "UpdateBookRequestDTO [uuid=" + uuid + ", title=" + title + ", author=" + author + ", genre=" + genre
        + ", image=" + image + "]";
  }
}
