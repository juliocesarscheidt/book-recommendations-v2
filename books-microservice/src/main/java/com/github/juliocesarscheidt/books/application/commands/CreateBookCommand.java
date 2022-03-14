package com.github.juliocesarscheidt.books.application.commands;

import com.github.juliocesarscheidt.books.application.service.BookService;

public class CreateBookCommand extends CommandBase {

  public CreateBookCommand() {
    this.commandName = "create";
    this.commandClass = BookService.class;
  }
}
