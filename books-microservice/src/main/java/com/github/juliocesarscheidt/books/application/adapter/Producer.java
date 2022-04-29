package com.github.juliocesarscheidt.books.application.adapter;

public interface Producer {

	public void sendMessage(String exchange, String replyQueueName, String messageString, Integer priority);
}
