package com.github.juliocesarscheidt.books.application.commands;

import com.github.juliocesarscheidt.books.application.service.BookService;

public class GetBookCommand extends CommandBase {

  public GetBookCommand() {
    this.commandName = "get";
    this.commandClass = BookService.class;
  }
}
