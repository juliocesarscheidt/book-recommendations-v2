package com.github.juliocesarscheidt.books.application.dto;

public class ListBookRequestDTO extends BaseDTO {
	
	public Integer page;
	
	public Integer size;

	public Integer getPage() {
		return page;
	}

	public void setPage(Integer page) {
		this.page = page;
	}

	public Integer getSize() {
		return size;
	}

	public void setSize(Integer size) {
		this.size = size;
	}

	@Override
	public String toString() {
		return "ListBookRequestDTO [page=" + page + ", size=" + size + "]";
	}	
}
