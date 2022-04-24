package com.github.juliocesarscheidt.books.infra.adapter;

import java.nio.charset.StandardCharsets;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessageDeliveryMode;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.amqp.core.MessageProperties;

import com.github.juliocesarscheidt.books.domain.adapter.Producer;

@Service
public class ProducerImpl implements Producer {

  @Autowired
  RabbitTemplate rabbitTemplate;

  public void sendMessage(String exchange, String replyQueueName, String messageString, Integer priority) {
    MessageProperties props = new MessageProperties();

    props.setContentType("application/json");
    props.setDeliveryMode(MessageDeliveryMode.PERSISTENT);
    if (priority != null) {
      props.setPriority(priority);
    }

    Message message = new Message(messageString.getBytes(StandardCharsets.UTF_8), props);
    rabbitTemplate.send(exchange, replyQueueName, message);
  }
}
