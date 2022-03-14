package com.github.juliocesarscheidt.books.application.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.dto.BaseDTO;
import com.github.juliocesarscheidt.books.application.dto.CreateBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.DeleteBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.GetBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.ListBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.UpdateBookRequestDTO;
import com.github.juliocesarscheidt.books.application.usecase.CreateBook;
import com.github.juliocesarscheidt.books.application.usecase.DeleteBook;
import com.github.juliocesarscheidt.books.application.usecase.GetBook;
import com.github.juliocesarscheidt.books.application.usecase.ListBook;
import com.github.juliocesarscheidt.books.application.usecase.UpdateBook;
import com.google.gson.Gson;

@Service
public class BookService extends BaseService {

  @Autowired
  CreateBook createBook;

  @Autowired
  GetBook getBook;

  @Autowired
  UpdateBook updateBook;

  @Autowired
  DeleteBook deleteBook;

  @Autowired
  ListBook listBook;

  public BaseDTO create(String messageDecoded) {
    Gson gson = new Gson();
    CreateBookRequestDTO requestDTO = gson.fromJson(messageDecoded, CreateBookRequestDTO.class);
    return createBook.execute(requestDTO);
  }

  public BaseDTO get(String messageDecoded) {
    Gson gson = new Gson();
    GetBookRequestDTO requestDTO = gson.fromJson(messageDecoded, GetBookRequestDTO.class);
    return getBook.execute(requestDTO);
  }

  public BaseDTO update(String messageDecoded) {
    Gson gson = new Gson();
    UpdateBookRequestDTO requestDTO = gson.fromJson(messageDecoded, UpdateBookRequestDTO.class);
    return updateBook.execute(requestDTO);
  }

  public BaseDTO delete(String messageDecoded) {
    Gson gson = new Gson();
    DeleteBookRequestDTO requestDTO = gson.fromJson(messageDecoded, DeleteBookRequestDTO.class);
    return deleteBook.execute(requestDTO);
  }

  public BaseDTO list(String messageDecoded) {
    Gson gson = new Gson();
    ListBookRequestDTO requestDTO = gson.fromJson(messageDecoded, ListBookRequestDTO.class);
    return listBook.execute(requestDTO);
  }
}
