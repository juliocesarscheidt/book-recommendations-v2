package com.github.juliocesarscheidt.books.application.usecase;

import java.util.ArrayList;
import java.util.List;

import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.dto.GetBookResponseDTO;
import com.github.juliocesarscheidt.books.application.dto.ListBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.ListBookResponseDTO;
import com.github.juliocesarscheidt.books.domain.entity.Book;
import com.github.juliocesarscheidt.books.infra.repository.BookRepository;

@Service
public class ListBook extends BaseUseCase {

  @Autowired
  BookRepository repository;

  public ListBookResponseDTO execute(ListBookRequestDTO dto) {
    ListBookResponseDTO response = new ListBookResponseDTO();
    List<GetBookResponseDTO> booksDto = new ArrayList<GetBookResponseDTO>();

    try {
      Pageable pageable = PageRequest.of(dto.getPage(), dto.getSize());
      Page<Book> books = repository.findAll(pageable);

      books.stream().forEach(book -> {
        GetBookResponseDTO dtoResponse = new GetBookResponseDTO();
        if (book != null) {
          BeanUtils.copyProperties(book, dtoResponse);
          booksDto.add(dtoResponse);
        }
      });

      response.setBooks(booksDto);
      return response;

    } catch (Exception e) {
      e.printStackTrace();
      return null;
    }
  }
}
