package com.github.juliocesarscheidt.books.infra.factory;

import org.springframework.beans.BeanUtils;

import com.github.juliocesarscheidt.books.application.dto.CreateBookRequestDTO;
import com.github.juliocesarscheidt.books.application.dto.UpdateBookRequestDTO;
import com.github.juliocesarscheidt.books.common.DatetimeUtils;
import com.github.juliocesarscheidt.books.common.StringUtils;
import com.github.juliocesarscheidt.books.domain.entity.Book;

public class BookFactory {

	public Book createBook(Book book, CreateBookRequestDTO dto) {
    BeanUtils.copyProperties(dto, book);

    book.setUuid(StringUtils.generateUuid(24));
    book.setCreatedAt(DatetimeUtils.getCurrentTimestampUTC());
    book.setUpdatedAt(DatetimeUtils.getCurrentTimestampUTC());

    return book;
	}

	public Book updateBook(Book book, UpdateBookRequestDTO dto) {
		if (dto.getTitle() != null) book.setTitle(dto.getTitle());
		if (dto.getAuthor() != null) book.setAuthor(dto.getAuthor());
		if (dto.getGenre() != null) book.setGenre(dto.getGenre());
		if (dto.getImage() != null) book.setImage(dto.getImage());
		book.setUpdatedAt(DatetimeUtils.getCurrentTimestampUTC());

		return book;
	}
}
