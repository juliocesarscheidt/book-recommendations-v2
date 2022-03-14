package com.github.juliocesarscheidt.books.application.commands;

import com.github.juliocesarscheidt.books.application.service.BaseService;

public abstract class CommandBase {
  public String commandName;
  public Class<? extends BaseService> commandClass;

  public String getCommandName() {
    return commandName;
  }
  public void setCommandName(String commandName) {
    this.commandName = commandName;
  }
  public Class<? extends BaseService> getCommandClass() {
    return commandClass;
  }
  public void setCommandClass(Class<? extends BaseService> commandClass) {
    this.commandClass = commandClass;
  }
}
