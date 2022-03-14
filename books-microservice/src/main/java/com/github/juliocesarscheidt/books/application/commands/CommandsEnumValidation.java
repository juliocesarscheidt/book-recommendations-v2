package com.github.juliocesarscheidt.books.application.commands;

public enum CommandsEnumValidation {
  // books
  CREATE_BOOK("books.create_book"),
  GET_BOOK("books.get_book"),
  UPDATE_BOOK("books.update_book"),
  DELETE_BOOK("books.delete_book"),
  LIST_BOOK("books.list_book"),
  // rates
  CREATE_RATE("rates.create_rate"),
  GET_RATE("rates.get_rate"),
  LIST_RATE("rates.list_rate");

  private String command;

  private CommandsEnumValidation(String command) {
    this.command = command;
  }

  public String getCommand() {
    return command;
  }
}
