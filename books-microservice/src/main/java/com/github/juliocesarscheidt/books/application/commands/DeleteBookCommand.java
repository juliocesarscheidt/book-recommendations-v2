package com.github.juliocesarscheidt.books.application.commands;

import com.github.juliocesarscheidt.books.application.service.BookService;

public class DeleteBookCommand extends CommandBase {

  public DeleteBookCommand() {
    this.commandName = "delete";
    this.commandClass = BookService.class;
  }
}
