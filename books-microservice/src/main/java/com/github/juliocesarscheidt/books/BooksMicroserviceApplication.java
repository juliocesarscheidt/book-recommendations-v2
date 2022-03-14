package com.github.juliocesarscheidt.books;

import org.springframework.amqp.rabbit.annotation.EnableRabbit;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@EnableRabbit
@SpringBootApplication
public class BooksMicroserviceApplication {

  public static void main(String[] args) {
    SpringApplication.run(BooksMicroserviceApplication.class, args);
  }
}
