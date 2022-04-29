package com.github.juliocesarscheidt.books.application.adapter;

import com.github.juliocesarscheidt.books.application.dto.BaseDTO;

public interface CommandHandler {

	public BaseDTO process(String commandString, String messageDecoded);
}
