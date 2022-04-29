package com.github.juliocesarscheidt.books.infra.adapter;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.amqp.core.MessageProperties;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.github.juliocesarscheidt.books.application.adapter.CommandHandler;
import com.github.juliocesarscheidt.books.application.adapter.Consumer;
import com.github.juliocesarscheidt.books.application.adapter.Producer;
import com.github.juliocesarscheidt.books.application.dto.BaseDTO;

@Component
public class ConsumerImpl implements Consumer {

  @Autowired
  Producer producer;

  @Autowired
  CommandHandler commandHandler;

  public BaseDTO processMessage(String messageDecoded, String routingKey) {
    System.out.println("RoutingKey :: " + routingKey);
    return commandHandler.process(routingKey, messageDecoded);
  }

  public void sendResponse(BaseDTO response, MessageProperties props) {
    Gson gson = new GsonBuilder()
        .setDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'")
        .create();
    String messageString = gson.toJson(response);
    producer.sendMessage("", props.getReplyTo(), messageString, null);
  }

  @RabbitListener(queues = "${amqp.books_queue}")
  public void consume(Message message) {
    try {
      String messageDecoded = new String(message.getBody(), "UTF-8");
      MessageProperties props = message.getMessageProperties();
      BaseDTO response = processMessage(messageDecoded, props.getReceivedRoutingKey());
      System.out.println("Response :: " + response);

      if (props.getReplyTo() != "") {
        sendResponse(response, props);
      }

    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
