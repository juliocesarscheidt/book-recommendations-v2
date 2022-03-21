package com.github.juliocesarscheidt.books.application.commands;

public enum CommandsEnumValidation {

  CREATE_BOOK("books.create_book"),
  GET_BOOK("books.get_book"),
  UPDATE_BOOK("books.update_book"),
  DELETE_BOOK("books.delete_book"),
  LIST_BOOK("books.list_book");

  private String command;

  private CommandsEnumValidation(String command) {
    this.command = command;
  }

  public String getCommand() {
    return command;
  }
}
