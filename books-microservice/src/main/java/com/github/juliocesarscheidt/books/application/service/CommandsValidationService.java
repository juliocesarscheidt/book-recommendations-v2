package com.github.juliocesarscheidt.books.application.service;

import java.util.HashMap;
import java.util.Map;

import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.commands.CommandBase;
import com.github.juliocesarscheidt.books.application.commands.CreateBookCommand;
import com.github.juliocesarscheidt.books.application.commands.DeleteBookCommand;
import com.github.juliocesarscheidt.books.application.commands.GetBookCommand;
import com.github.juliocesarscheidt.books.application.commands.ListBookCommand;
import com.github.juliocesarscheidt.books.application.commands.UpdateBookCommand;

@Service
public class CommandsValidationService {

  private Map<String, CommandBase> commandsMap = new HashMap<>();

  public CommandsValidationService() {
    // books
    this.commandsMap.put("books.create_book", new CreateBookCommand());
    this.commandsMap.put("books.get_book", new GetBookCommand());
    this.commandsMap.put("books.update_book", new UpdateBookCommand());
    this.commandsMap.put("books.delete_book", new DeleteBookCommand());
    this.commandsMap.put("books.list_book", new ListBookCommand());
  }

  public CommandBase getCommand(String command) {
    try {
      return this.commandsMap.get(command);

    } catch (Exception e) {
      System.out.println(e.getMessage());
      return null;
    }
  }
}
