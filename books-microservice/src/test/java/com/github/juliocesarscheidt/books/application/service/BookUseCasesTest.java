package com.github.juliocesarscheidt.books.application.service;

import java.util.HashMap;
import java.util.Map;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.junit4.SpringRunner;

import com.github.juliocesarscheidt.books.application.dto.CreateBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.CreateBookResponseDTO;
import com.github.juliocesarscheidt.books.application.dto.DeleteBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.GetBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.GetBookResponseDTO;
import com.github.juliocesarscheidt.books.application.dto.ListBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.ListBookResponseDTO;
import com.github.juliocesarscheidt.books.application.dto.UpdateBookRequestDTO;
import com.github.juliocesarscheidt.books.application.usecase.CreateBook;
import com.github.juliocesarscheidt.books.application.usecase.DeleteBook;
import com.github.juliocesarscheidt.books.application.usecase.GetBook;
import com.github.juliocesarscheidt.books.application.usecase.ListBook;
import com.github.juliocesarscheidt.books.application.usecase.UpdateBook;
import com.github.juliocesarscheidt.books.infra.repository.BookRepository;
import com.google.gson.Gson;

@RunWith(SpringRunner.class)
@SpringBootTest
@ActiveProfiles("test")
public class BookUseCasesTest {

  @Mock
  BookRepository repository;

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

  public Map<String, Object> createBook() {
    Gson gson = new Gson();

    String createBookRequest = "{\"title\":\"Clean Architecture\",\"author\":\"Robert Martin\",\"genre\":\"Software\",\"image\":\"https://fakeimage.com/image.jpg\"}";
    CreateBookRequestDTO createBookRequestDTO = gson.fromJson(createBookRequest, CreateBookRequestDTO.class);
    CreateBookResponseDTO createBookResponseDTO = createBook.execute(createBookRequestDTO);

    Map<String, Object> response = new HashMap<String, Object>();
    response.put("requestDto", createBookRequestDTO);
    response.put("responseDto", createBookResponseDTO);

    return response;
  }

  @Test
  void createBookTest() {
    Map<String, Object> createBookMap = this.createBook();
    CreateBookResponseDTO createBookResponseDTO = (CreateBookResponseDTO) createBookMap.get("responseDto");
    GetBookRequestDTO getBookRequestDTO = new GetBookRequestDTO();
    getBookRequestDTO.setUuid(createBookResponseDTO.getUuid());
    GetBookResponseDTO getBookResponseDTO = getBook.execute(getBookRequestDTO);

    Assertions.assertEquals(getBookResponseDTO.getTitle(), (((CreateBookRequestDTO) createBookMap.get("requestDto")).getTitle()));
    Assertions.assertEquals(getBookResponseDTO.getAuthor(), (((CreateBookRequestDTO) createBookMap.get("requestDto")).getAuthor()));
    Assertions.assertEquals(getBookResponseDTO.getGenre(), (((CreateBookRequestDTO) createBookMap.get("requestDto")).getGenre()));
    Assertions.assertEquals(getBookResponseDTO.getImage(), (((CreateBookRequestDTO) createBookMap.get("requestDto")).getImage()));

    // clean up
    DeleteBookRequestDTO deleteBookRequestDTO = new DeleteBookRequestDTO();
    deleteBookRequestDTO.setUuid(createBookResponseDTO.getUuid());
    deleteBook.execute(deleteBookRequestDTO);
  }

  @Test
  void getBookTest() {
    GetBookRequestDTO getBookRequestDTO = new GetBookRequestDTO();
    getBookRequestDTO.setUuid(null);
    GetBookResponseDTO getBookResponseDTO = getBook.execute(getBookRequestDTO);
    Assertions.assertEquals(getBookResponseDTO, null);
  }

  @Test
  void updateBookTest() {
    Map<String, Object> createBookMap = this.createBook();
    CreateBookResponseDTO createBookResponseDTO = (CreateBookResponseDTO) createBookMap.get("responseDto");
    GetBookRequestDTO getBookRequestDTO = new GetBookRequestDTO();
    getBookRequestDTO.setUuid(createBookResponseDTO.getUuid());
    GetBookResponseDTO getBookResponseDTO = getBook.execute(getBookRequestDTO);

    Assertions.assertEquals(getBookResponseDTO.getUuid(), createBookResponseDTO.getUuid());

    UpdateBookRequestDTO updateBookRequestDTO = new UpdateBookRequestDTO();
    updateBookRequestDTO.setUuid(createBookResponseDTO.getUuid());
    updateBookRequestDTO.setTitle("Clean Architecture - Revised Edition");
    updateBook.execute(updateBookRequestDTO);

    getBookResponseDTO = getBook.execute(getBookRequestDTO);

    Assertions.assertNotEquals(getBookResponseDTO.getTitle(), "Clean Architecture");

    // clean up
    DeleteBookRequestDTO deleteBookRequestDTO = new DeleteBookRequestDTO();
    deleteBookRequestDTO.setUuid(createBookResponseDTO.getUuid());
    deleteBook.execute(deleteBookRequestDTO);
  }

  @Test
  void deleteBookTest() {
    Map<String, Object> createBookMap = this.createBook();
    CreateBookResponseDTO createBookResponseDTO = (CreateBookResponseDTO) createBookMap.get("responseDto");
    GetBookRequestDTO getBookRequestDTO = new GetBookRequestDTO();
    getBookRequestDTO.setUuid(createBookResponseDTO.getUuid());
    GetBookResponseDTO getBookResponseDTO = getBook.execute(getBookRequestDTO);

    Assertions.assertEquals(getBookResponseDTO.getUuid(), createBookResponseDTO.getUuid());

    DeleteBookRequestDTO deleteBookRequestDTO = new DeleteBookRequestDTO();
    deleteBookRequestDTO.setUuid(createBookResponseDTO.getUuid());
    deleteBook.execute(deleteBookRequestDTO);

    getBookResponseDTO = getBook.execute(getBookRequestDTO);
    Assertions.assertEquals(getBookResponseDTO, null);
  }

  @Test
  void listBookTestPage0WithoutContent() {
    ListBookRequestDTO listBookRequestDTO = new ListBookRequestDTO();
    listBookRequestDTO.setPage(0);
    listBookRequestDTO.setSize(50);
    ListBookResponseDTO listBookResponseDTO = listBook.execute(listBookRequestDTO);

    Assertions.assertEquals(listBookResponseDTO.getBooks().size(), 0);
  }

  @Test
  void listBookTestPage0WithContentShouldReturnSomething() {
    Map<String, Object> createBookMap = this.createBook();
    CreateBookResponseDTO createBookResponseDTO = (CreateBookResponseDTO) createBookMap.get("responseDto");
    GetBookRequestDTO getBookRequestDTO = new GetBookRequestDTO();
    getBookRequestDTO.setUuid(createBookResponseDTO.getUuid());
    GetBookResponseDTO getBookResponseDTO = getBook.execute(getBookRequestDTO);

    ListBookRequestDTO listBookRequestDTO = new ListBookRequestDTO();
    listBookRequestDTO.setPage(0);
    listBookRequestDTO.setSize(50);
    ListBookResponseDTO listBookResponseDTO = listBook.execute(listBookRequestDTO);

    Assertions.assertEquals(listBookResponseDTO.getBooks().size(), 1);

      // clean up
    DeleteBookRequestDTO deleteBookRequestDTO = new DeleteBookRequestDTO();
    deleteBookRequestDTO.setUuid(getBookResponseDTO.getUuid());
    deleteBook.execute(deleteBookRequestDTO);
  }

  @Test
    void listBookTestPage1WithContentShouldReturnNothing() {
    Map<String, Object> createBookMap = this.createBook();
    CreateBookResponseDTO createBookResponseDTO = (CreateBookResponseDTO) createBookMap.get("responseDto");
    GetBookRequestDTO getBookRequestDTO = new GetBookRequestDTO();
    getBookRequestDTO.setUuid(createBookResponseDTO.getUuid());
    GetBookResponseDTO getBookResponseDTO = getBook.execute(getBookRequestDTO);

    ListBookRequestDTO listBookRequestDTO = new ListBookRequestDTO();
    listBookRequestDTO.setPage(1);
    listBookRequestDTO.setSize(50);
    ListBookResponseDTO listBookResponseDTO = listBook.execute(listBookRequestDTO);

    Assertions.assertEquals(listBookResponseDTO.getBooks().size(), 0);

    // clean up
    DeleteBookRequestDTO deleteBookRequestDTO = new DeleteBookRequestDTO();
    deleteBookRequestDTO.setUuid(getBookResponseDTO.getUuid());
    deleteBook.execute(deleteBookRequestDTO);
  }
}
