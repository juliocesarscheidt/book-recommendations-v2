package com.github.juliocesarscheidt.books.application.dto;

import java.util.Date;

import com.github.juliocesarscheidt.books.domain.service.DatetimeConverter;
import com.google.gson.annotations.SerializedName;

public class GetBookResponseDTO extends BaseDTO {

  public String uuid;

  public String title;

  public String author;

  public String genre;

  public String image;

  @SerializedName("created_at")
  public Date createdAt;

  @SerializedName("updated_at")
  public Date updatedAt;

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

  public Date getCreatedAt() {
    return DatetimeConverter.convertDateToGMT(this.createdAt);
  }

  public void setCreatedAt(Date createdAt) {
    this.createdAt = createdAt;
  }

  public Date getUpdatedAt() {
    return DatetimeConverter.convertDateToGMT(this.updatedAt);
  }

  public void setUpdatedAt(Date updatedAt) {
    this.updatedAt = updatedAt;
  }

  @Override
  public String toString() {
    return "GetBookResponseDTO [uuid=" + uuid + ", title=" + title + ", author=" + author + ", genre=" + genre
        + ", image=" + image + ", createdAt=" + createdAt + ", updatedAt=" + updatedAt + "]";
  }
}
