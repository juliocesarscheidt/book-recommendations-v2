package com.github.juliocesarscheidt.books.infra.adapter;

import java.lang.reflect.Method;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.github.juliocesarscheidt.books.application.adapter.CommandHandler;
import com.github.juliocesarscheidt.books.application.commands.CommandBase;
import com.github.juliocesarscheidt.books.application.dto.BaseDTO;
import com.github.juliocesarscheidt.books.application.service.BaseService;
import com.github.juliocesarscheidt.books.application.service.BookService;
import com.github.juliocesarscheidt.books.application.service.CommandsValidationService;

@Service
public class CommandHandlerImpl implements CommandHandler {

  @Autowired
  BookService bookService;

  @Autowired
  CommandsValidationService commandsValidationService;

  public BaseDTO process(String commandString, String messageDecoded) {
    try {
      CommandBase cmdInstance = commandsValidationService.getCommand(commandString);
      if (cmdInstance.equals(null)) {
        return null;
      }

      Class<? extends BaseService> commandClass = cmdInstance.getCommandClass();
      String commandName = cmdInstance.getCommandName();
      Method method = commandClass.getDeclaredMethod(commandName, String.class);

      return (BaseDTO) method.invoke(this.bookService, messageDecoded);

    } catch (Exception e) {
      e.printStackTrace();
      return null;
    }
  }
}
