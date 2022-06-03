package com.github.juliocesarscheidt.books.application.usecase;

import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.dto.GetBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.GetBookResponseDTO;
import com.github.juliocesarscheidt.books.domain.entity.Book;
import com.github.juliocesarscheidt.books.infra.repository.BookRepository;

@Service
public class GetBook {

  @Autowired
  BookRepository repository;

  public GetBookResponseDTO execute(GetBookRequestDTO dto) {
    if (dto.getUuid() == null) return null;

    try {
      Book book = repository.findByUuid(dto.getUuid());
      if (book == null) return null;

      GetBookResponseDTO response = new GetBookResponseDTO();
      BeanUtils.copyProperties(book, response);
      return response;

    } catch (Exception e) {
      e.printStackTrace();
      return null;
    }
  }
}
