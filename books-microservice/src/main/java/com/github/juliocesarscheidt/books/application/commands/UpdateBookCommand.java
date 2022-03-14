package com.github.juliocesarscheidt.books.application.commands;

import com.github.juliocesarscheidt.books.application.service.BookService;

public class UpdateBookCommand extends CommandBase {

  public UpdateBookCommand() {
    this.commandName = "update";
    this.commandClass = BookService.class;
  }
}
