package com.github.juliocesarscheidt.books.application.adapter;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessageProperties;

import com.github.juliocesarscheidt.books.application.dto.BaseDTO;

public interface Consumer {

	public BaseDTO processMessage(String messageDecoded, String routingKey);

	public void sendResponse(BaseDTO response, MessageProperties props);

	public void consume(Message message);
}
