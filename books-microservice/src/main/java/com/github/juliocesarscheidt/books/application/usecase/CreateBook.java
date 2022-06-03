package com.github.juliocesarscheidt.books.application.usecase;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.dto.CreateBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.CreateBookResponseDTO;
import com.github.juliocesarscheidt.books.domain.entity.Book;
import com.github.juliocesarscheidt.books.infra.factory.BookFactory;
import com.github.juliocesarscheidt.books.infra.repository.BookRepository;

@Service
public class CreateBook {

  @Autowired
  BookRepository repository;

  public CreateBookResponseDTO execute(CreateBookRequestDTO dto) {
    BookFactory factory = new BookFactory();

    try {
      Book book = new Book();
      book = factory.createBook(book, dto);

      repository.save(book);
      return new CreateBookResponseDTO(book.uuid);

    } catch (Exception e) {
      e.printStackTrace();
      return null;
    }
  }
}
