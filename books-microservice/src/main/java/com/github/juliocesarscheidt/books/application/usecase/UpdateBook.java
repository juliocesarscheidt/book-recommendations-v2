package com.github.juliocesarscheidt.books.application.usecase;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.dto.UpdateBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.UpdateBookResponseDTO;
import com.github.juliocesarscheidt.books.domain.entity.Book;
import com.github.juliocesarscheidt.books.infra.factory.BookFactory;
import com.github.juliocesarscheidt.books.infra.repository.BookRepository;

@Service
public class UpdateBook {

  @Autowired
  BookRepository repository;

  public UpdateBookResponseDTO execute(UpdateBookRequestDTO dto) {
    if (dto.getUuid() == null) return null;

    try {
      BookFactory factory = new BookFactory();
      Book book = repository.findByUuid(dto.getUuid());
      if (book == null) return null;

      book = factory.updateBook(book, dto);

      repository.save(book);
      return new UpdateBookResponseDTO();

    } catch (Exception e) {
      e.printStackTrace();
      return null;
    }
  }
}
