package com.github.juliocesarscheidt.books.config;

import org.springframework.amqp.rabbit.connection.CachingConnectionFactory;
import org.springframework.amqp.rabbit.core.RabbitAdmin;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class RabbitConfiguration {

  @Value("${amqp.conn_string}")
  String amqpConnString;

  @Bean
  public CachingConnectionFactory connectionFactory() {
      System.out.println("amqpConnString :: " + amqpConnString);
      CachingConnectionFactory rabbitConnectionFactory = new CachingConnectionFactory();
      rabbitConnectionFactory.setUri(amqpConnString);
      rabbitConnectionFactory.setVirtualHost("/");
      return rabbitConnectionFactory;
  }

  @Bean
  public RabbitAdmin amqpAdmin() {
	  return new RabbitAdmin(connectionFactory());
  }

  @Bean
  public RabbitTemplate rabbitTemplate() {
	  return new RabbitTemplate(connectionFactory());
  }
}
