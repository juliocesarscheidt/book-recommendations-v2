package com.github.juliocesarscheidt.books.application.commands;

import com.github.juliocesarscheidt.books.application.service.BookService;

public class ListBookCommand extends CommandBase {

  public ListBookCommand() {
    this.commandName = "list";
    this.commandClass = BookService.class;
  }
}
