package com.github.juliocesarscheidt.books.infra.repository;

import javax.transaction.Transactional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import com.github.juliocesarscheidt.books.domain.entity.Book;

@Repository
public interface BookRepository extends JpaRepository<Book, String> {

	@Query(value = "SELECT b FROM book b WHERE uuid = :uuid", nativeQuery = false)
	Book findByUuid(@Param("uuid") String uuid);

	@Transactional
    @Modifying
	@Query(value = "DELETE book b WHERE uuid = :uuid", nativeQuery = false)
	void deleteByUuid(@Param("uuid") String uuid);
}
