package com.github.juliocesarscheidt.books.config;

public enum RabbitPriorities {

  LOWER_PRIORITY(1),
  MEDIUM_PRIORITY(2),
  HIGHER_PRIORITY(3);

  Integer priority;

  RabbitPriorities(int priority) {
    this.setPriority(priority);
  }

  public Integer getPriority() {
    return priority;
  }

  public void setPriority(Integer priority) {
    this.priority = priority;
  }
}
