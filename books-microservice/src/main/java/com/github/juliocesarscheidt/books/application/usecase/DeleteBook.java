package com.github.juliocesarscheidt.books.application.usecase;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.dto.DeleteBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.DeleteBookResponseDTO;
import com.github.juliocesarscheidt.books.domain.entity.Book;
import com.github.juliocesarscheidt.books.infra.repository.BookRepository;

@Service
public class DeleteBook extends BaseUseCase {

  @Autowired
  BookRepository repository;

  public DeleteBookResponseDTO execute(DeleteBookRequestDTO dto) {
	if (dto.getUuid() == null) return null;

    try {
      Book book = repository.findByUuid(dto.getUuid());
      if (book == null) return null;

      repository.deleteByUuid(dto.uuid);
      return new DeleteBookResponseDTO();

    } catch (Exception e) {
      e.printStackTrace();
      return null;
    }
  }
}
