package com.github.juliocesarscheidt.books.application.usecase;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.dto.UpdateBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.UpdateBookResponseDTO;
import com.github.juliocesarscheidt.books.domain.entity.Book;
import com.github.juliocesarscheidt.books.infra.repository.BookRepository;

@Service
public class UpdateBook extends BaseUseCase {

  @Autowired
  BookRepository repository;;

  public UpdateBookResponseDTO execute(UpdateBookRequestDTO dto) {
    if (dto.getUuid() == null) return null;

    try {
	  Book book = repository.findByUuid(dto.getUuid());
	  if (book == null) return null;

	  if (dto.getTitle() != null) book.setTitle(dto.getTitle());
	  if (dto.getAuthor() != null) book.setAuthor(dto.getAuthor());
	  if (dto.getGenre() != null) book.setGenre(dto.getGenre());
	  if (dto.getImage() != null) book.setImage(dto.getImage());
	  book.setUpdatedAt(getCurrentTimestampUTC());

      repository.save(book);
      return new UpdateBookResponseDTO();

    } catch (Exception e) {
      e.printStackTrace();
      return null;
    }
  }
}
