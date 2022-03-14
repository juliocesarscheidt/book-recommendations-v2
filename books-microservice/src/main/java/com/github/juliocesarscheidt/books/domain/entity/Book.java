package com.github.juliocesarscheidt.books.domain.entity;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;

import com.github.juliocesarscheidt.books.domain.service.DatetimeConverter;

@Entity(name = "book")
public class Book {

  @Id
  @Column(name = "uuid", length = 24, nullable = false)
  public String uuid;

  @Column(name = "title", length = 255, nullable = false)
  public String title;

  @Column(name = "author", length = 255, nullable = false)
  public String author;

  @Column(name = "genre", length = 255, nullable = false)
  public String genre;

  @Column(name = "image", length = 255, nullable = true)
  public String image;

  @Column(name = "created_at")
  public Date createdAt;

  @Column(name = "updated_at")
  public Date updatedAt;

  public Book() {
  }

  public Book(String uuid, String title, String author, String genre, String image, Date createdAt, Date updatedAt) {
    this.uuid = uuid;
    this.title = title;
    this.author = author;
    this.genre = genre;
    this.image = image;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
  }

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
    return "Book [uuid=" + uuid + ", title=" + title + ", author=" + author + ", genre=" + genre + ", image=" + image
        + ", createdAt=" + createdAt + ", updatedAt=" + updatedAt + "]";
  }
}
