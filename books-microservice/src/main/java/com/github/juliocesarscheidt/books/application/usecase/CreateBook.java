package com.github.juliocesarscheidt.books.application.usecase;

import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.dto.CreateBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.CreateBookResponseDTO;
import com.github.juliocesarscheidt.books.domain.entity.Book;
import com.github.juliocesarscheidt.books.infra.repository.BookRepository;

@Service
public class CreateBook extends BaseUseCase {

  @Autowired
  BookRepository repository;
  
  public CreateBookResponseDTO execute(CreateBookRequestDTO dto) {
    Book book = new Book();
    BeanUtils.copyProperties(dto, book);

	book.setUuid(generateUuid());
    book.setCreatedAt(getCurrentTimestampUTC());
    book.setUpdatedAt(getCurrentTimestampUTC());
 
    try {
    	repository.save(book);

        return new CreateBookResponseDTO(book.uuid);

      } catch (Exception e) {
        e.printStackTrace();
        return null;
      }

  }
}
